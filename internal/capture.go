package internal

import (
	"fmt"
	"sync"

	"github.com/gordonklaus/portaudio"
)

const (
	SampleRate   = 16000
	Channels     = 1
	FrameDurMs   = 20
	FrameSamples = SampleRate * FrameDurMs / 1000 // 320 samples per 20ms frame
)

type Capture struct {
	stream *portaudio.Stream
	frames chan []int16
	stop   chan struct{}
	wg     sync.WaitGroup
}

// NewCapture creates a mic capture that delivers 20ms PCM16 frames on a buffered channel.
func NewCapture(bufferSize int) (*Capture, error) {
	if bufferSize <= 0 {
		bufferSize = 64
	}
	c := &Capture{
		frames: make(chan []int16, bufferSize),
		stop:   make(chan struct{}),
	}
	return c, nil
}

// Frames returns the channel delivering audio frames.
func (c *Capture) Frames() <-chan []int16 {
	return c.frames
}

// Start opens the default mic and begins capturing.
func (c *Capture) Start() error {
	buf := make([]int16, FrameSamples)

	stream, err := portaudio.OpenDefaultStream(Channels, 0, float64(SampleRate), FrameSamples, buf)
	if err != nil {
		return fmt.Errorf("open mic stream: %w", err)
	}
	c.stream = stream

	if err := stream.Start(); err != nil {
		stream.Close()
		return fmt.Errorf("start mic stream: %w", err)
	}

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			select {
			case <-c.stop:
				return
			default:
			}
			if err := stream.Read(); err != nil {
				// Stream closed or error — exit
				return
			}
			frame := make([]int16, FrameSamples)
			copy(frame, buf)
			select {
			case c.frames <- frame:
			default:
				// Drop frame if consumer is too slow
			}
		}
	}()

	return nil
}

// Stop halts capture and releases resources.
func (c *Capture) Stop() {
	close(c.stop)
	c.wg.Wait()
	if c.stream != nil {
		c.stream.Stop()
		c.stream.Close()
	}
	close(c.frames)
}
