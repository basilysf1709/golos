<img width="1024" height="448" alt="Screenshot 1447-09-06 at 3 18 35 AM" src="https://github.com/user-attachments/assets/62e4b2ce-28a5-4edc-a62f-abd055606061" />


# Golos

Golos is an extremely lightweight Wispr Flow alternative. Hold a hotkey, speak, release, and your words get pasted.

## Install

```bash
curl -fsSL https://raw.githubusercontent.com/basilysf1709/golos/main/install.sh | bash
```

## Permissions

Golos needs **Accessibility** permission to listen for your hotkey and paste transcriptions. After running Golos for the first time, macOS will prompt you to grant access:

1. Open **System Settings → Privacy & Security → Accessibility**
2. Enable the toggle for your terminal app (e.g. Terminal, iTerm2, Alacritty)

If you skip this step, Golos won't be able to detect the hotkey or paste text.

## Usage

```bash
golos setup                  # configure API key
golos                        # run in foreground
golos -d                     # run in background
golos --output stdout        # output to stdout instead of clipboard
golos --hotkey cmd           # override hotkey
golos stop                   # stop background process
```

## CLI Commands

| Command | Description |
|---------|-------------|
| `golos setup` | Configure Deepgram API key |
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
