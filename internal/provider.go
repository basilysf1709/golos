package internal

// TranscriptResult holds a transcript from the STT provider.
type TranscriptResult struct {
	Text    string
	IsFinal bool
	SpeechFinal bool
}

// Provider is the interface for speech-to-text backends.
type Provider interface {
	Connect() error
	Write(p []byte) (int, error)
	Results() <-chan TranscriptResult
	Finalize() error
	Close()
}
