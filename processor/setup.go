package processor

import (
	"flag"
	"fmt"
	"os"

	"github.com/gordonklaus/portaudio"

	"github.com/iqbalyusuf/golos/internal"
)

type App struct {
	Proc   *Processor
	Hotkey internal.HotkeyInfo
	Config *internal.Config
}

func applyFlags(cfg *internal.Config, output, hotkey *string) {
	if *output != "" {
		cfg.OutputMode = *output
	}
	if *hotkey != "" {
		cfg.Hotkey = *hotkey
	}
}

func resolveOutput(cfg *internal.Config) internal.OutputMode {
	switch cfg.OutputMode {
	case "stdout":
		return &internal.StdoutMode{}
	case "clipboard":
		return &internal.ClipboardMode{}
	default:
		return nil
	}
}

func Setup() (*App, error) {
	outputFlag := flag.String("output", "", "output mode: clipboard or stdout (default: from config)")
	hotkeyFlag := flag.String("hotkey", "", "push-to-talk hotkey (default: from config)")
	flag.Parse()

	cfg, err := internal.LoadConfig()
	if err != nil {
		return nil, err
	}

	applyFlags(cfg, outputFlag, hotkeyFlag)

	out := resolveOutput(cfg)
	if out == nil {
		return nil, fmt.Errorf("unknown output mode: %s", cfg.OutputMode)
	}

	// Check accessibility permission for clipboard mode
	if cfg.OutputMode == "clipboard" {
		if !internal.CheckAccessibility() {
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprintln(os.Stderr, "  Accessibility permission required!")
			fmt.Fprintln(os.Stderr, "  Go to: System Settings → Privacy & Security → Accessibility")
			fmt.Fprintln(os.Stderr, "  Add your terminal app (Terminal, iTerm2, etc.) to the list.")
			fmt.Fprintln(os.Stderr, "")
			os.Exit(1)
		}
	}

	// Resolve hotkey
	hk, err := internal.ResolveHotkey(cfg.Hotkey)
	if err != nil {
		return nil, err
	}

	// Initialize PortAudio
	if err := portaudio.Initialize(); err != nil {
		return nil, fmt.Errorf("PortAudio init: %w", err)
	}

	// Create processor
	proc, err := New(cfg, out)
	if err != nil {
		portaudio.Terminate()
		return nil, err
	}

	return &App{Proc: proc, Hotkey: hk, Config: cfg}, nil
}
