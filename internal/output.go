package internal

// OutputMode defines how transcribed text is delivered.
type OutputMode interface {
	// Deliver outputs the transcribed text.
	Deliver(text string) error
}
