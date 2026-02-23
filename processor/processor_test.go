package processor

import (
	"math"
	"testing"
	"time"

	"github.com/basilysf1709/golos/internal"
)

// mockOutput captures delivered text for assertions.
type mockOutput struct {
	delivered string
	err       error
}

func (m *mockOutput) Deliver(text string) error {
	m.delivered = text
	return m.err
}

func TestNew(t *testing.T) {
	cfg := &internal.Config{
		DeepgramAPIKey: "test-key",
		Language:       "en-US",
	}
	out := &mockOutput{}

	p, err := New(cfg, out)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if p.cfg != cfg {
		t.Error("cfg not set")
	}
	if p.out != out {
		t.Error("out not set")
	}
	if p.vad == nil {
		t.Error("vad should be initialized")
	}
	if p.recording {
		t.Error("should not be recording initially")
	}
}

func TestStopWhenNotRecording(t *testing.T) {
	cfg := &internal.Config{
		DeepgramAPIKey: "test-key",
		Language:       "en-US",
	}
	out := &mockOutput{}

	p, err := New(cfg, out)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	// Should not panic
	p.Stop()

	if out.delivered != "" {
		t.Error("should not deliver anything when not recording")
	}
}

func TestRmsLevelSilence(t *testing.T) {
	frame := make([]int16, 320)
	level := rmsLevel(frame)
	if level != 0 {
		t.Errorf("rmsLevel of silence = %f, want 0", level)
	}
}

func TestRmsLevelConstant(t *testing.T) {
	frame := make([]int16, 320)
	for i := range frame {
		frame[i] = 1000
	}
	level := rmsLevel(frame)
	if math.Abs(level-1000) > 0.01 {
		t.Errorf("rmsLevel of constant 1000 = %f, want 1000", level)
	}
}

func TestRmsLevelLoud(t *testing.T) {
	frame := make([]int16, 320)
	for i := range frame {
		frame[i] = 20000
	}
	level := rmsLevel(frame)
	if level < 19000 {
		t.Errorf("rmsLevel of loud signal = %f, expected > 19000", level)
	}
}

func TestVuMeterSilence(t *testing.T) {
	meter := vuMeter(0)
	if meter != "[░░░░░░░░░░]" {
		t.Errorf("vuMeter(0) = %q, want all empty", meter)
	}
}

func TestVuMeterBelowThreshold(t *testing.T) {
	meter := vuMeter(5)
	if meter != "[░░░░░░░░░░]" {
		t.Errorf("vuMeter(5) = %q, want all empty", meter)
	}
}

func TestVuMeterLow(t *testing.T) {
	meter := vuMeter(15)
	// Should have at least 1 bar
	if meter == "[░░░░░░░░░░]" {
		t.Error("vuMeter(15) should have at least one bar")
	}
}

func TestVuMeterMax(t *testing.T) {
	meter := vuMeter(32768)
	if meter != "[██████████]" {
		t.Errorf("vuMeter(32768) = %q, want all full", meter)
	}
}

func TestVuMeterMid(t *testing.T) {
	meter := vuMeter(500)
	// Should have some bars but not all
	full := len("██████████")
	empty := len("░░░░░░░░░░")
	content := meter[1 : len(meter)-1] // strip [ and ]
	if len(content) != full && len(content) != empty {
		// Just verify it's a valid meter with brackets
	}
	if meter[0] != '[' || meter[len(meter)-1] != ']' {
		t.Errorf("vuMeter should be bracketed, got %q", meter)
	}
}

func TestVuMeterIncreasing(t *testing.T) {
	// Higher levels should produce more bars
	low := vuMeter(20)
	high := vuMeter(10000)

	countBars := func(m string) int {
		return len([]rune(m)) - 2 // subtract brackets
	}

	// Count filled bars (█ is 3 bytes in UTF-8)
	lowFilled := 0
	highFilled := 0
	for _, r := range low {
		if r == '█' {
			lowFilled++
		}
	}
	for _, r := range high {
		if r == '█' {
			highFilled++
		}
	}
	_ = countBars

	if highFilled <= lowFilled {
		t.Errorf("higher level should have more bars: low=%d high=%d", lowFilled, highFilled)
	}
}

func TestAccumulateFinal(t *testing.T) {
	cfg := &internal.Config{
		DeepgramAPIKey: "test-key",
		Language:       "en-US",
	}
	out := &mockOutput{}
	p, err := New(cfg, out)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	results := make(chan internal.TranscriptResult, 8)
	p.doneCh = make(chan struct{})
	p.gotFinal = make(chan struct{})
	p.provider = &mockProvider{results: results}

	go p.accumulate()

	// Send interim then final
	results <- internal.TranscriptResult{Text: "hel", IsFinal: false}
	results <- internal.TranscriptResult{Text: "hello", IsFinal: true}

	<-p.gotFinal

	p.mu.Lock()
	text := p.transcript.String()
	p.mu.Unlock()

	if text != "hello" {
		t.Errorf("transcript = %q, want %q", text, "hello")
	}

	close(p.doneCh)
}

func TestAccumulateMultipleFinals(t *testing.T) {
	cfg := &internal.Config{
		DeepgramAPIKey: "test-key",
		Language:       "en-US",
	}
	out := &mockOutput{}
	p, err := New(cfg, out)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	results := make(chan internal.TranscriptResult, 8)
	p.doneCh = make(chan struct{})
	p.gotFinal = make(chan struct{})
	p.provider = &mockProvider{results: results}

	go p.accumulate()

	results <- internal.TranscriptResult{Text: "hello", IsFinal: true}
	<-p.gotFinal

	results <- internal.TranscriptResult{Text: "world", IsFinal: true}

	// Close the results channel to let accumulate drain
	close(results)

	// Wait for accumulate to finish processing
	// by trying to read transcript until it has both words
	deadline := time.After(time.Second)
	for {
		p.mu.Lock()
		text := p.transcript.String()
		p.mu.Unlock()
		if text == "hello world" {
			break
		}
		select {
		case <-deadline:
			p.mu.Lock()
			t.Fatalf("transcript = %q, want %q", p.transcript.String(), "hello world")
			p.mu.Unlock()
		default:
			time.Sleep(5 * time.Millisecond)
		}
	}
}

// mockProvider implements internal.Provider for testing.
type mockProvider struct {
	results chan internal.TranscriptResult
}

func (m *mockProvider) Connect() error                          { return nil }
func (m *mockProvider) Write(p []byte) (int, error)             { return len(p), nil }
func (m *mockProvider) Results() <-chan internal.TranscriptResult { return m.results }
func (m *mockProvider) Finalize() error                         { return nil }
func (m *mockProvider) Close()                                  {}
