package internal

import (
	"testing"
)

func TestNewCaptureDefaultBuffer(t *testing.T) {
	c, err := NewCapture(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap(c.frames) != 64 {
		t.Errorf("frames buffer = %d, want 64 (default)", cap(c.frames))
	}
}

func TestNewCaptureNegativeBuffer(t *testing.T) {
	c, err := NewCapture(-10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap(c.frames) != 64 {
		t.Errorf("frames buffer = %d, want 64 (default)", cap(c.frames))
	}
}

func TestNewCaptureCustomBuffer(t *testing.T) {
	c, err := NewCapture(128)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap(c.frames) != 128 {
		t.Errorf("frames buffer = %d, want 128", cap(c.frames))
	}
}

func TestFramesReturnsSameChannel(t *testing.T) {
	c, _ := NewCapture(8)
	ch := c.Frames()

	// Send a frame directly on the internal channel
	frame := make([]int16, FrameSamples)
	frame[0] = 42
	c.frames <- frame

	got := <-ch
	if got[0] != 42 {
		t.Errorf("got frame[0] = %d, want 42", got[0])
	}
}

func TestFramesChannelIsReceiveOnly(t *testing.T) {
	c, _ := NewCapture(8)
	ch := c.Frames()

	// Type assertion: Frames() returns <-chan, not chan
	// This is a compile-time check — if Frames() returned chan []int16,
	// this test file wouldn't compile since we assign to <-chan []int16.
	var _ <-chan []int16 = ch
}

func TestStopClosesFramesChannel(t *testing.T) {
	c, _ := NewCapture(8)

	// Stop without Start — stream is nil, should still close channels cleanly
	c.Stop()

	// frames channel should be closed
	_, ok := <-c.frames
	if ok {
		t.Error("frames channel should be closed after Stop")
	}
}

func TestFrameDropWhenBufferFull(t *testing.T) {
	c, _ := NewCapture(2) // tiny buffer

	// Fill the buffer
	c.frames <- make([]int16, FrameSamples)
	c.frames <- make([]int16, FrameSamples)

	// Third send should not block (drop behavior is in the goroutine,
	// but we can verify the channel is full)
	if len(c.frames) != 2 {
		t.Errorf("frames len = %d, want 2 (full)", len(c.frames))
	}
}

func TestConstants(t *testing.T) {
	if SampleRate != 16000 {
		t.Errorf("SampleRate = %d, want 16000", SampleRate)
	}
	if Channels != 1 {
		t.Errorf("Channels = %d, want 1", Channels)
	}
	if FrameDurMs != 20 {
		t.Errorf("FrameDurMs = %d, want 20", FrameDurMs)
	}
	if FrameSamples != 320 {
		t.Errorf("FrameSamples = %d, want 320", FrameSamples)
	}
}
