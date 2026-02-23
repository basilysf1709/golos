<img width="1024" height="448" alt="image" src="https://github.com/user-attachments/assets/3da18627-624e-4e5c-bbe0-e9a433db51c9" />


# Golos

Golos is an extremely lightweight Wispr Flow alternative. Hold a hotkey, speak, release, and your words get pasted.

## Install

```bash
curl -fsSL https://raw.githubusercontent.com/basilysf1709/golos/main/install.sh | bash
```

## Usage

```bash
golos                        # run in foreground
golos -d                     # run in background
golos --output stdout        # output to stdout instead of clipboard
golos --hotkey cmd           # override hotkey
golos stop                   # stop background process
```

## CLI Commands

| Command | Description |
|---------|-------------|
| `golos` | Run speech-to-text (foreground) |
| `golos -d` | Run speech-to-text (background) |
| `golos stop` | Stop the background process |
| `golos add <phrase> <replacement>` | Add a dictionary replacement |
| `golos delete <phrase>` | Delete a dictionary entry |
| `golos list` | List all dictionary entries |
| `golos import <file.toml>` | Import dictionary from a TOML file |

### Flags

| Flag | Description |
|------|-------------|
| `-d`, `--detach` | Run in background |
| `--output <mode>` | Override output mode (`clipboard` or `stdout`) |
| `--hotkey <key>` | Override hotkey |

### Dictionary

Manage word/phrase replacements that are applied to transcriptions:

```bash
golos add "period" "."
golos add "new line" "\n"
golos delete "period"
golos list
golos import dictionary.example.toml
```

## Configuration

Config file: `~/.config/golos/config.toml`

```toml
deepgram_api_key = "your-key"
hotkey = "right_option"
output_mode = "clipboard"
sample_rate = 16000
language = "en-US"
```

Environment variables `DEEPGRAM_API_KEY`, `GOLOS_OUTPUT`, and `GOLOS_HOTKEY` override config values.

## Requirements

macOS, a [Deepgram API key](https://console.deepgram.com), and Accessibility permission for your terminal.
