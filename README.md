<p align="center">
  <img src="assets/mascot.png" alt="Golos Mascot" width="200" />
</p>

<h1 align="center">Golos</h1>

<p align="center"><strong>Voice-to-Text, Instantly</strong></p>

<p align="center">
An extremely lightweight Wispr Flow alternative. Hold a hotkey, speak, release, and your words get pasted.<br/>
One binary. Zero bloat. Speech that actually works for you.
</p>

<p align="center">
  <a href="https://golos.sh">Website</a> &bull;
  <a href="#install">Quick Start</a> &bull;
  <a href="https://github.com/basilysf1709/golos">GitHub</a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/platform-macOS-blue" alt="platform" />
  <img src="https://img.shields.io/badge/license-MIT-green" alt="license" />
  <img src="https://img.shields.io/github/v/release/basilysf1709/golos?color=orange&label=version" alt="version" />
  <img src="https://img.shields.io/badge/language-Go-00ADD8" alt="language" />
</p>

## What is Golos?

Golos (Russian for "voice") is a free, open-source macOS CLI that turns your voice into text wherever your cursor is. Hold a hotkey to record, release to transcribe, and the result is instantly pasted into the focused application.

**How it works:**

1. **Hold** your hotkey (default: `Right Option`) to start recording
2. **Speak** — audio streams to [Deepgram Nova-3](https://deepgram.com) in real time with a live transcript in your terminal
3. **Release** — the final transcription is pasted into whatever app is focused via `Cmd+V`

**Key features:**

- **Push-to-talk** — no always-on microphone, only records while the hotkey is held
- **Single binary** — no Electron, no GUI, no background service required
- **Foreground or background mode** — run interactively or daemonize with `golos -d`
- **Live feedback** — VU meter and interim transcript displayed in real time while speaking
- **Dictionary replacements** — map spoken words to text (e.g. say "period" → `.`, "new line" → `\n`)
- **Configurable hotkey** — `right_option`, `right_command`, `fn`, `f18`, or `f19`
- **Two output modes** — paste into focused app (`clipboard`) or print to `stdout` for piping
- **Config layering** — defaults → config file → environment variables → CLI flags

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
