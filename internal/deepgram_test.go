package internal

import (
	"testing"

	api "github.com/deepgram/deepgram-go-sdk/v3/pkg/api/listen/v1/websocket/interfaces"
)

func TestNewDeepgram(t *testing.T) {
	d := NewDeepgram("test-key", "en-US")
	if d.apiKey != "test-key" {
		t.Errorf("apiKey = %q, want %q", d.apiKey, "test-key")
	}
	if d.lang != "en-US" {
		t.Errorf("lang = %q, want %q", d.lang, "en-US")
	}
	if d.results == nil {
		t.Error("results channel should not be nil")
	}
	if cap(d.results) != 64 {
		t.Errorf("results buffer = %d, want 64", cap(d.results))
	}
	if d.ctx == nil {
		t.Error("ctx should not be nil")
	}
}

func TestDeepgramImplementsProvider(t *testing.T) {
	var _ Provider = NewDeepgram("key", "en")
}

func TestWriteBeforeConnect(t *testing.T) {
	d := NewDeepgram("key", "en")
	_, err := d.Write([]byte("audio"))
	if err == nil {
		t.Error("expected error writing before Connect")
	}
}

func TestFinalizeBeforeConnect(t *testing.T) {
	d := NewDeepgram("key", "en")
	err := d.Finalize()
	if err != nil {
		t.Errorf("Finalize before Connect should return nil, got: %v", err)
	}
}

func TestCloseBeforeConnect(t *testing.T) {
	d := NewDeepgram("key", "en")
	// Should not panic
	d.Close()
}

func TestResultsChannel(t *testing.T) {
	d := NewDeepgram("key", "en")
	ch := d.Results()
	if ch == nil {
		t.Fatal("Results() returned nil")
	}

	// Verify it's the same channel
	d.results <- TranscriptResult{Text: "hello", IsFinal: true}
	got := <-ch
	if got.Text != "hello" || !got.IsFinal {
		t.Errorf("got %+v, want {Text:hello IsFinal:true}", got)
	}
}

// --- deepgramCallback tests ---

func newTestCallback(bufSize int) (*deepgramCallback, chan TranscriptResult) {
	ch := make(chan TranscriptResult, bufSize)
	return &deepgramCallback{results: ch}, ch
}

func TestCallbackMessageFinal(t *testing.T) {
	cb, ch := newTestCallback(8)

	mr := &api.MessageResponse{
		IsFinal:     true,
		SpeechFinal: false,
	}
	mr.Channel.Alternatives = []api.Alternative{
		{Transcript: "hello world"},
	}

	if err := cb.Message(mr); err != nil {
		t.Fatalf("Message error: %v", err)
	}

	got := <-ch
	if got.Text != "hello world" {
		t.Errorf("Text = %q, want %q", got.Text, "hello world")
	}
	if !got.IsFinal {
		t.Error("IsFinal should be true")
	}
}

func TestCallbackMessageInterim(t *testing.T) {
	cb, ch := newTestCallback(8)

	mr := &api.MessageResponse{
		IsFinal:     false,
		SpeechFinal: false,
	}
	mr.Channel.Alternatives = []api.Alternative{
		{Transcript: "hel"},
	}

	_ = cb.Message(mr)

	got := <-ch
	if got.Text != "hel" {
		t.Errorf("Text = %q, want %q", got.Text, "hel")
	}
	if got.IsFinal {
		t.Error("IsFinal should be false for interim")
	}
}

func TestCallbackMessageEmptyAlternatives(t *testing.T) {
	cb, ch := newTestCallback(8)

	mr := &api.MessageResponse{}
	mr.Channel.Alternatives = []api.Alternative{}

	_ = cb.Message(mr)

	if len(ch) != 0 {
		t.Error("empty alternatives should produce no result")
	}
}

func TestCallbackMessageBlankTranscript(t *testing.T) {
	cb, ch := newTestCallback(8)

	mr := &api.MessageResponse{}
	mr.Channel.Alternatives = []api.Alternative{
		{Transcript: "   "},
	}

	_ = cb.Message(mr)

	if len(ch) != 0 {
		t.Error("blank transcript should produce no result")
	}
}

func TestCallbackMessageTrimsWhitespace(t *testing.T) {
	cb, ch := newTestCallback(8)

	mr := &api.MessageResponse{IsFinal: true}
	mr.Channel.Alternatives = []api.Alternative{
		{Transcript: "  hello  "},
	}

	_ = cb.Message(mr)

	got := <-ch
	if got.Text != "hello" {
		t.Errorf("Text = %q, want %q (trimmed)", got.Text, "hello")
	}
}

func TestCallbackMessageDropsWhenFull(t *testing.T) {
	cb, ch := newTestCallback(1) // buffer of 1

	// Fill the buffer
	mr := &api.MessageResponse{IsFinal: true}
	mr.Channel.Alternatives = []api.Alternative{
		{Transcript: "first"},
	}
	_ = cb.Message(mr)

	// This should be dropped, not block
	mr2 := &api.MessageResponse{IsFinal: true}
	mr2.Channel.Alternatives = []api.Alternative{
		{Transcript: "second"},
	}
	_ = cb.Message(mr2)

	if len(ch) != 1 {
		t.Errorf("channel len = %d, want 1 (second message dropped)", len(ch))
	}
	got := <-ch
	if got.Text != "first" {
		t.Errorf("Text = %q, want %q", got.Text, "first")
	}
}

func TestCallbackSpeechFinalFlag(t *testing.T) {
	cb, ch := newTestCallback(8)

	mr := &api.MessageResponse{
		IsFinal:     true,
		SpeechFinal: true,
	}
	mr.Channel.Alternatives = []api.Alternative{
		{Transcript: "done"},
	}

	_ = cb.Message(mr)

	got := <-ch
	if !got.SpeechFinal {
		t.Error("SpeechFinal should be true")
	}
}

func TestCallbackStubsReturnNil(t *testing.T) {
	cb, _ := newTestCallback(1)

	if err := cb.Open(nil); err != nil {
		t.Errorf("Open: %v", err)
	}
	if err := cb.Metadata(nil); err != nil {
		t.Errorf("Metadata: %v", err)
	}
	if err := cb.SpeechStarted(nil); err != nil {
		t.Errorf("SpeechStarted: %v", err)
	}
	if err := cb.UtteranceEnd(nil); err != nil {
		t.Errorf("UtteranceEnd: %v", err)
	}
	if err := cb.Close(nil); err != nil {
		t.Errorf("Close: %v", err)
	}
	if err := cb.UnhandledEvent(nil); err != nil {
		t.Errorf("UnhandledEvent: %v", err)
	}
	if err := cb.Error(&api.ErrorResponse{ErrCode: "test", Description: "test err"}); err != nil {
		t.Errorf("Error: %v", err)
	}
}
