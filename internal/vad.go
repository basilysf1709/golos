package internal

import (
	"encoding/binary"

	webrtcvad "github.com/maxhawkins/go-webrtcvad"
)

type VADEvent int

const (
	VADNone        VADEvent = iota
	SpeechStart
	SpeechEnd
)

type Detector struct {
	vad            *webrtcvad.VAD
	sampleRate     int
	hangoverFrames int
	active         bool
	silentCount    int
}

// NewDetector creates a VAD with the given hangover duration in milliseconds.
// mode: 0 (least aggressive) to 3 (most aggressive).
func NewDetector(sampleRate, hangoverMs, mode int) (*Detector, error) {
	v, err := webrtcvad.New()
	if err != nil {
		return nil, err
	}
	if err := v.SetMode(mode); err != nil {
		return nil, err
	}

	frameDurMs := 20
	hangoverFrames := hangoverMs / frameDurMs
	if hangoverFrames < 1 {
		hangoverFrames = 1
	}

	return &Detector{
		vad:            v,
		sampleRate:     sampleRate,
		hangoverFrames: hangoverFrames,
	}, nil
}

// Process takes a 20ms frame of int16 samples and returns the VAD event.
func (d *Detector) Process(frame []int16) (VADEvent, error) {
	// Convert int16 slice to bytes for go-webrtcvad
	buf := make([]byte, len(frame)*2)
	for i, s := range frame {
		binary.LittleEndian.PutUint16(buf[i*2:], uint16(s))
	}

	active, err := d.vad.Process(d.sampleRate, buf)
	if err != nil {
		return VADNone, err
	}

	if active {
		d.silentCount = 0
		if !d.active {
			d.active = true
			return SpeechStart, nil
		}
		return VADNone, nil
	}

	// Frame is silent
	if d.active {
		d.silentCount++
		if d.silentCount >= d.hangoverFrames {
			d.active = false
			d.silentCount = 0
			return SpeechEnd, nil
		}
	}
	return VADNone, nil
}

// IsActive returns whether speech is currently detected.
func (d *Detector) IsActive() bool {
	return d.active
}

// Reset resets the VAD state.
func (d *Detector) Reset() {
	d.active = false
	d.silentCount = 0
}
