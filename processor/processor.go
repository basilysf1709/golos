package processor

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/basilysf1709/golos/internal"
)

type Processor struct {
	cfg        *internal.Config
	out        internal.OutputMode
	dict       *internal.Dictionary
	vad        *internal.Detector
	mu         sync.Mutex
	recording  bool
	capture    *internal.Capture
	provider   internal.Provider
	transcript strings.Builder
	doneCh     chan struct{}
	gotFinal   chan struct{}
	connected  chan struct{} // closed when Deepgram connection is ready
	streamWg   sync.WaitGroup // ensures streamAudio() finishes before Finalize()
}

func New(cfg *internal.Config, out internal.OutputMode) (*Processor, error) {
	vad, err := internal.NewDetector(internal.SampleRate, 300, 3)
	if err != nil {
		return nil, fmt.Errorf("VAD init: %w", err)
	}
	return &Processor{cfg: cfg, out: out, vad: vad, dict: internal.LoadDictionary()}, nil
}

func (p *Processor) Start() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.recording {
		return
	}
	p.recording = true
	p.transcript.Reset()
	p.vad.Reset()

	fmt.Print("\r\033[K🎙  Listening...")
	internal.OverlayShow(0)

	// Start mic immediately — no waiting for network
	var err error
	p.capture, err = internal.NewCapture(128)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nMic error: %v\n", err)
		p.recording = false
		return
	}
	if err := p.capture.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "\nMic start error: %v\n", err)
		p.recording = false
		return
	}

	p.doneCh = make(chan struct{})
	p.gotFinal = make(chan struct{})
	p.connected = make(chan struct{})

	// Connect to Deepgram in background
	go p.connect()

	// Audio → STT (buffers until connected)
	p.streamWg.Add(1)
	go p.streamAudio()

	// Transcript accumulator
	go p.accumulate()
}

func (p *Processor) connect() {
	prov := internal.NewDeepgram(p.cfg.DeepgramAPIKey, p.cfg.Language)
	if err := prov.Connect(); err != nil {
		fmt.Fprintf(os.Stderr, "\nSTT connect error: %v\n", err)
		p.mu.Lock()
		p.recording = false
		p.mu.Unlock()
		return
	}
	p.mu.Lock()
	p.provider = prov
	p.mu.Unlock()
	close(p.connected)
}

func (p *Processor) Stop() {
	p.mu.Lock()
	if !p.recording {
		p.mu.Unlock()
		return
	}
	p.recording = false
	internal.OverlayShow(1)

	cap := p.capture
	prov := p.provider
	done := p.doneCh
	gf := p.gotFinal
	conn := p.connected
	p.mu.Unlock()

	// Wait for connection if still pending (with timeout)
	select {
	case <-conn:
		// Connection finished — re-capture provider since it may have been
		// nil when we first read it (connect() assigns it under lock).
		p.mu.Lock()
		prov = p.provider
		p.mu.Unlock()
	case <-time.After(2 * time.Second):
	}

	// 1. Stop mic — closes the frames channel so streamAudio() drains
	//    any remaining buffered frames and exits.
	if cap != nil {
		cap.Stop()
	}

	// 2. Wait for streamAudio() to finish sending all frames to Deepgram.
	//    Without this, Finalize() could fire before the last frames are written.
	p.streamWg.Wait()

	// 3. Tell Deepgram we're done sending audio.
	if prov != nil {
		_ = prov.Finalize()
	}

	// 4. Wait for at least one final result, then drain remaining results.
	select {
	case <-gf:
	case <-time.After(2 * time.Second):
	}
	// Drain any results still in the channel after the first final.
	p.drainResults(prov, 300*time.Millisecond)

	// 5. Now safe to tear down.
	if done != nil {
		close(done)
	}
	if prov != nil {
		prov.Close()
	}

	p.mu.Lock()
	finalText := p.transcript.String()
	p.mu.Unlock()

	if finalText != "" {
		finalText = p.dict.Replace(finalText)
		fmt.Print("\r\033[K")
		if err := p.out.Deliver(finalText); err != nil {
			fmt.Fprintf(os.Stderr, "Output error: %v\n", err)
		}
		fmt.Print("\r\033[K")
	} else {
		fmt.Print("\r\033[K(no speech detected)\n")
	}
	internal.OverlayHide()
}

// drainResults reads any remaining results from the provider channel
// until no more arrive within the timeout window.
func (p *Processor) drainResults(prov internal.Provider, timeout time.Duration) {
	if prov == nil {
		return
	}
	for {
		select {
		case result, ok := <-prov.Results():
			if !ok {
				return
			}
			p.mu.Lock()
			if result.IsFinal {
				if p.transcript.Len() > 0 {
					p.transcript.WriteString(" ")
				}
				p.transcript.WriteString(result.Text)
				fmt.Printf("\r\033[K💬 %s", p.transcript.String())
			}
			p.mu.Unlock()
		case <-time.After(timeout):
			return
		}
	}
}

func (p *Processor) streamAudio() {
	defer p.streamWg.Done()

	// Buffer frames while Deepgram connects
	var buffered [][]byte
	var prov internal.Provider

	for frame := range p.capture.Frames() {
		_, _ = p.vad.Process(frame)

		level := rmsLevel(frame)
		meter := vuMeter(level)
		fmt.Printf("\r\033[K🎙  Listening %s", meter)

		buf := make([]byte, len(frame)*2)
		for i, s := range frame {
			binary.LittleEndian.PutUint16(buf[i*2:], uint16(s))
		}

		select {
		case <-p.connected:
			// Capture provider once after connection is ready
			if prov == nil {
				p.mu.Lock()
				prov = p.provider
				p.mu.Unlock()
			}
			if prov == nil {
				continue
			}
			// Connected — flush buffer then stream normally
			if len(buffered) > 0 {
				for _, b := range buffered {
					_, _ = prov.Write(b)
				}
				buffered = nil
			}
			_, _ = prov.Write(buf)
		default:
			// Still connecting — buffer the audio
			buffered = append(buffered, buf)
		}
	}
}

func (p *Processor) accumulate() {
	// Wait for connection before reading results
	select {
	case <-p.connected:
	case <-p.doneCh:
		return
	}

	finalSignaled := false
	for {
		select {
		case <-p.doneCh:
			return
		case result, ok := <-p.provider.Results():
			if !ok {
				return
			}
			p.mu.Lock()
			if result.IsFinal {
				if p.transcript.Len() > 0 {
					p.transcript.WriteString(" ")
				}
				p.transcript.WriteString(result.Text)
				fmt.Printf("\r\033[K💬 %s", p.transcript.String())
				if !finalSignaled {
					finalSignaled = true
					close(p.gotFinal)
				}
			} else {
				interim := p.transcript.String()
				if interim != "" {
					interim += " "
				}
				fmt.Printf("\r\033[K💬 %s%s", interim, result.Text)
			}
			p.mu.Unlock()
		}
	}
}

func rmsLevel(frame []int16) float64 {
	var sum float64
	for _, s := range frame {
		sum += float64(s) * float64(s)
	}
	return math.Sqrt(sum / float64(len(frame)))
}

func vuMeter(level float64) string {
	if level < 10 {
		return "[░░░░░░░░░░]"
	}
	logLevel := math.Log2(level)
	bars := int((logLevel - 3) * 10 / 12)
	if bars < 1 {
		bars = 1
	}
	if bars > 10 {
		bars = 10
	}
	return "[" + strings.Repeat("█", bars) + strings.Repeat("░", 10-bars) + "]"
}
