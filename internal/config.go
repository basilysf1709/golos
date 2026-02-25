package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	DeepgramAPIKey string `toml:"deepgram_api_key"`
	Hotkey         string `toml:"hotkey"`
	OutputMode     string `toml:"output_mode"`
	SampleRate     int    `toml:"sample_rate"`
	Language       string `toml:"language"`
	Overlay        bool   `toml:"overlay"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Hotkey:     "right_option",
		OutputMode: "clipboard",
		SampleRate: 16000,
		Language:   "en-US",
		Overlay:    true,
	}

	// Load .env file from current directory (silent if missing)
	_ = godotenv.Load()

	// Try config file
	if home, err := os.UserHomeDir(); err == nil {
		configPath := filepath.Join(home, ".config", "golos", "config.toml")
		if _, err := os.Stat(configPath); err == nil {
			if _, err := toml.DecodeFile(configPath, cfg); err != nil {
				return nil, fmt.Errorf("parsing config file: %w", err)
			}
		}
	}

	// Env vars override config file
	if key := os.Getenv("DEEPGRAM_API_KEY"); key != "" {
		cfg.DeepgramAPIKey = key
	}
	if mode := os.Getenv("GOLOS_OUTPUT"); mode != "" {
		cfg.OutputMode = mode
	}
	if hotkey := os.Getenv("GOLOS_HOTKEY"); hotkey != "" {
		cfg.Hotkey = hotkey
	}

	if cfg.DeepgramAPIKey == "" {
		return nil, fmt.Errorf("DEEPGRAM_API_KEY is required (set via env var or ~/.config/golos/config.toml)")
	}

	return cfg, nil
}
