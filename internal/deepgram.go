package internal

import (
	"context"
	"fmt"
	"strings"
	"sync"

	api "github.com/deepgram/deepgram-go-sdk/v3/pkg/api/listen/v1/websocket/interfaces"
	interfaces "github.com/deepgram/deepgram-go-sdk/v3/pkg/client/interfaces"
	client "github.com/deepgram/deepgram-go-sdk/v3/pkg/client/listen"
)

var initOnce sync.Once

func initSDK() {
	initOnce.Do(func() {
		client.Init(client.InitLib{
			LogLevel: client.LogLevelDefault,
		})
	})
}

type DeepgramProvider struct {
	apiKey  string
	lang    string
	results chan TranscriptResult
	dgConn  *client.WSCallback
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewDeepgram(apiKey, language string) *DeepgramProvider {
	ctx, cancel := context.WithCancel(context.Background())
	return &DeepgramProvider{
		apiKey:  apiKey,
		lang:    language,
		results: make(chan TranscriptResult, 64),
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (d *DeepgramProvider) Connect() error {
	initSDK()

	cOptions := &interfaces.ClientOptions{
		EnableKeepAlive: true,
	}
	if d.apiKey != "" {
		cOptions.APIKey = d.apiKey
	}

	tOptions := &interfaces.LiveTranscriptionOptions{
		Model:          "nova-3",
		Language:       d.lang,
		Punctuate:      true,
		Encoding:       "linear16",
		Channels:       1,
		SampleRate:     16000,
		SmartFormat:    true,
		InterimResults: true,
		VadEvents:      true,
		Endpointing:    "300",
		UtteranceEndMs: "1000",
	}

	cb := &deepgramCallback{results: d.results}

	conn, err := client.NewWSUsingCallback(d.ctx, "", cOptions, tOptions, cb)
	if err != nil {
		return fmt.Errorf("deepgram connection: %w", err)
	}

	if !conn.Connect() {
		return fmt.Errorf("deepgram websocket connect failed")
	}

	d.dgConn = conn
	return nil
}

func (d *DeepgramProvider) Write(p []byte) (int, error) {
	if d.dgConn == nil {
		return 0, fmt.Errorf("not connected")
	}
	return d.dgConn.Write(p)
}

func (d *DeepgramProvider) Results() <-chan TranscriptResult {
	return d.results
}

func (d *DeepgramProvider) Finalize() error {
	if d.dgConn != nil {
		return d.dgConn.Finalize()
	}
	return nil
}

func (d *DeepgramProvider) Close() {
	if d.dgConn != nil {
		d.dgConn.Stop()
	}
	d.cancel()
}

// deepgramCallback implements the LiveMessageCallback interface.
type deepgramCallback struct {
	results chan TranscriptResult
}

func (c *deepgramCallback) Open(_ *api.OpenResponse) error {
	return nil
}

func (c *deepgramCallback) Message(mr *api.MessageResponse) error {
	if len(mr.Channel.Alternatives) == 0 {
		fmt.Println("\n[DEBUG callback] Message: no alternatives")
		return nil
	}
	text := strings.TrimSpace(mr.Channel.Alternatives[0].Transcript)
	if text == "" {
		fmt.Println("\n[DEBUG callback] Message: empty transcript")
		return nil
	}

	fmt.Printf("\n[DEBUG callback] Message: text=%q isFinal=%v speechFinal=%v\n", text, mr.IsFinal, mr.SpeechFinal)

	result := TranscriptResult{
		Text:        text,
		IsFinal:     mr.IsFinal,
		SpeechFinal: mr.SpeechFinal,
	}

	select {
	case c.results <- result:
		fmt.Println("[DEBUG callback] result sent to channel")
	default:
		fmt.Println("[DEBUG callback] result DROPPED (channel full)")
	}
	return nil
}

func (c *deepgramCallback) Metadata(_ *api.MetadataResponse) error {
	return nil
}

func (c *deepgramCallback) SpeechStarted(_ *api.SpeechStartedResponse) error {
	return nil
}

func (c *deepgramCallback) UtteranceEnd(_ *api.UtteranceEndResponse) error {
	return nil
}

func (c *deepgramCallback) Close(_ *api.CloseResponse) error {
	return nil
}

func (c *deepgramCallback) Error(er *api.ErrorResponse) error {
	fmt.Printf("[Deepgram Error] %s: %s\n", er.ErrCode, er.Description)
	return nil
}

func (c *deepgramCallback) UnhandledEvent(_ []byte) error {
	return nil
}
