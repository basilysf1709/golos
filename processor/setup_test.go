package processor

import (
	"testing"

	"github.com/basilysf1709/golos/internal"
)

func TestResolveOutputStdout(t *testing.T) {
	cfg := &internal.Config{OutputMode: "stdout"}
	out := resolveOutput(cfg)
	if out == nil {
		t.Fatal("expected StdoutMode, got nil")
	}
	if _, ok := out.(*internal.StdoutMode); !ok {
		t.Errorf("expected *StdoutMode, got %T", out)
	}
}

func TestResolveOutputClipboard(t *testing.T) {
	cfg := &internal.Config{OutputMode: "clipboard"}
	out := resolveOutput(cfg)
	if out == nil {
		t.Fatal("expected ClipboardMode, got nil")
	}
	if _, ok := out.(*internal.ClipboardMode); !ok {
		t.Errorf("expected *ClipboardMode, got %T", out)
	}
}

func TestResolveOutputUnknown(t *testing.T) {
	cfg := &internal.Config{OutputMode: "fax"}
	out := resolveOutput(cfg)
	if out != nil {
		t.Errorf("expected nil for unknown mode, got %T", out)
	}
}

func TestApplyFlags(t *testing.T) {
	cfg := &internal.Config{
		OutputMode: "clipboard",
		Hotkey:     "right_option",
	}

	output := "stdout"
	hotkey := "fn"
	applyFlags(cfg, &output, &hotkey)

	if cfg.OutputMode != "stdout" {
		t.Errorf("OutputMode = %q, want %q", cfg.OutputMode, "stdout")
	}
	if cfg.Hotkey != "fn" {
		t.Errorf("Hotkey = %q, want %q", cfg.Hotkey, "fn")
	}
}

func TestApplyFlagsEmpty(t *testing.T) {
	cfg := &internal.Config{
		OutputMode: "clipboard",
		Hotkey:     "right_option",
	}

	output := ""
	hotkey := ""
	applyFlags(cfg, &output, &hotkey)

	if cfg.OutputMode != "clipboard" {
		t.Errorf("OutputMode should not change, got %q", cfg.OutputMode)
	}
	if cfg.Hotkey != "right_option" {
		t.Errorf("Hotkey should not change, got %q", cfg.Hotkey)
	}
}
