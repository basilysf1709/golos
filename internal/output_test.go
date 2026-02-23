package internal

import (
	"testing"
)

func TestStdoutModeImplementsOutputMode(t *testing.T) {
	var _ OutputMode = &StdoutMode{}
}

func TestClipboardModeImplementsOutputMode(t *testing.T) {
	var _ OutputMode = &ClipboardMode{}
}

func TestStdoutModeDeliver(t *testing.T) {
	s := &StdoutMode{}
	// Should not error
	if err := s.Deliver("hello"); err != nil {
		t.Errorf("Deliver error: %v", err)
	}
}
