package internal

import (
	"math"
	"testing"
)

const testSampleRate = 16000
const testFrameSize = 320 // 20ms at 16kHz

func silentFrame() []int16 {
	return make([]int16, testFrameSize)
}

func loudFrame() []int16 {
	frame := make([]int16, testFrameSize)
	for i := range frame {
		// Generate a 400Hz sine wave at high amplitude — reliably triggers VAD
		frame[i] = int16(20000 * math.Sin(2*math.Pi*400*float64(i)/float64(testSampleRate)))
	}
	return frame
}

func newTestDetector(t *testing.T, hangoverMs, mode int) *Detector {
	t.Helper()
	d, err := NewDetector(testSampleRate, hangoverMs, mode)
	if err != nil {
		t.Fatalf("NewDetector: %v", err)
	}
	return d
}

func TestNewDetector(t *testing.T) {
	d, err := NewDetector(testSampleRate, 300, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d.hangoverFrames != 15 {
		t.Errorf("hangoverFrames = %d, want 15", d.hangoverFrames)
	}
	if d.active {
		t.Error("new detector should not be active")
	}
}

func TestNewDetectorMinHangover(t *testing.T) {
	d, err := NewDetector(testSampleRate, 0, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d.hangoverFrames != 1 {
		t.Errorf("hangoverFrames = %d, want 1 (minimum)", d.hangoverFrames)
	}
}

func TestNewDetectorInvalidMode(t *testing.T) {
	_, err := NewDetector(testSampleRate, 300, 99)
	if err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestSilenceProducesNoEvents(t *testing.T) {
	d := newTestDetector(t, 300, 3)
	for i := 0; i < 50; i++ {
		ev, err := d.Process(silentFrame())
		if err != nil {
			t.Fatalf("frame %d: %v", i, err)
		}
		if ev != VADNone {
			t.Fatalf("frame %d: got %d, want VADNone", i, ev)
		}
	}
	if d.IsActive() {
		t.Error("detector should not be active after only silence")
	}
}

func TestSpeechStartOnVoice(t *testing.T) {
	d := newTestDetector(t, 300, 3)

	// Feed loud frames until we get SpeechStart
	gotStart := false
	for i := 0; i < 10; i++ {
		ev, err := d.Process(loudFrame())
		if err != nil {
			t.Fatalf("frame %d: %v", i, err)
		}
		if ev == SpeechStart {
			gotStart = true
			break
		}
	}
	if !gotStart {
		t.Error("expected SpeechStart from loud frames")
	}
	if !d.IsActive() {
		t.Error("detector should be active after SpeechStart")
	}
}

func TestNoDoubleSpeechStart(t *testing.T) {
	d := newTestDetector(t, 300, 3)

	startCount := 0
	for i := 0; i < 30; i++ {
		ev, err := d.Process(loudFrame())
		if err != nil {
			t.Fatalf("frame %d: %v", i, err)
		}
		if ev == SpeechStart {
			startCount++
		}
	}
	if startCount != 1 {
		t.Errorf("SpeechStart fired %d times, want 1", startCount)
	}
}

func TestSpeechEndAfterHangover(t *testing.T) {
	d := newTestDetector(t, 60, 3) // 60ms = 3 frames hangover

	// Trigger speech
	for i := 0; i < 10; i++ {
		_, _ = d.Process(loudFrame())
	}
	if !d.IsActive() {
		t.Fatal("detector should be active after loud frames")
	}

	// Feed silent frames — should get SpeechEnd after 3 frames
	gotEnd := false
	for i := 0; i < 10; i++ {
		ev, err := d.Process(silentFrame())
		if err != nil {
			t.Fatalf("frame %d: %v", i, err)
		}
		if ev == SpeechEnd {
			gotEnd = true
			break
		}
	}
	if !gotEnd {
		t.Error("expected SpeechEnd after hangover period")
	}
	if d.IsActive() {
		t.Error("detector should not be active after SpeechEnd")
	}
}

func TestHangoverPreventsEarlyEnd(t *testing.T) {
	d := newTestDetector(t, 100, 3) // 100ms = 5 frames hangover

	// Trigger speech
	for i := 0; i < 10; i++ {
		_, _ = d.Process(loudFrame())
	}

	// Feed only 2 silent frames — should NOT trigger SpeechEnd
	for i := 0; i < 2; i++ {
		ev, err := d.Process(silentFrame())
		if err != nil {
			t.Fatalf("frame %d: %v", i, err)
		}
		if ev == SpeechEnd {
			t.Fatal("SpeechEnd fired too early, hangover should prevent it")
		}
	}

	// Resume speech — should reset silence counter, no SpeechEnd
	for i := 0; i < 5; i++ {
		ev, err := d.Process(loudFrame())
		if err != nil {
			t.Fatalf("frame %d: %v", i, err)
		}
		if ev == SpeechEnd {
			t.Fatal("SpeechEnd should not fire after speech resumed")
		}
	}
	if !d.IsActive() {
		t.Error("detector should still be active")
	}
}

func TestReset(t *testing.T) {
	d := newTestDetector(t, 300, 3)

	// Trigger speech
	for i := 0; i < 10; i++ {
		_, _ = d.Process(loudFrame())
	}
	if !d.IsActive() {
		t.Fatal("detector should be active")
	}

	d.Reset()

	if d.IsActive() {
		t.Error("detector should not be active after Reset")
	}
	if d.silentCount != 0 {
		t.Errorf("silentCount = %d, want 0 after Reset", d.silentCount)
	}
}

func TestFullCycle(t *testing.T) {
	d := newTestDetector(t, 40, 3) // 40ms = 2 frames hangover

	var events []VADEvent

	// Silence
	for i := 0; i < 5; i++ {
		ev, _ := d.Process(silentFrame())
		if ev != VADNone {
			events = append(events, ev)
		}
	}

	// Speech
	for i := 0; i < 10; i++ {
		ev, _ := d.Process(loudFrame())
		if ev != VADNone {
			events = append(events, ev)
		}
	}

	// Silence until end
	for i := 0; i < 10; i++ {
		ev, _ := d.Process(silentFrame())
		if ev != VADNone {
			events = append(events, ev)
		}
	}

	if len(events) != 2 {
		t.Fatalf("got %d events, want 2 (SpeechStart + SpeechEnd)", len(events))
	}
	if events[0] != SpeechStart {
		t.Errorf("events[0] = %d, want SpeechStart", events[0])
	}
	if events[1] != SpeechEnd {
		t.Errorf("events[1] = %d, want SpeechEnd", events[1])
	}
}
