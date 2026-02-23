#!/bin/bash
set -e

REPO="iqbalyusuf/golos"
INSTALL_DIR="/usr/local/bin"
CONFIG_DIR="$HOME/.config/golos"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BOLD='\033[1m'
NC='\033[0m'

info()  { echo -e "${BOLD}$1${NC}"; }
ok()    { echo -e "${GREEN}✓${NC} $1"; }
warn()  { echo -e "${YELLOW}!${NC} $1"; }
fail()  { echo -e "${RED}✗${NC} $1"; exit 1; }

# ---

info "Installing golos — speech-to-text for Claude Code"
echo ""

# Check macOS
if [[ "$(uname)" != "Darwin" ]]; then
    fail "golos only supports macOS"
fi

# Check Homebrew
if ! command -v brew &>/dev/null; then
    fail "Homebrew is required. Install it from https://brew.sh"
fi
ok "Homebrew found"

# Check Go
if ! command -v go &>/dev/null; then
    info "Installing Go..."
    brew install go
fi
ok "Go $(go version | awk '{print $3}' | sed 's/go//')"

# Install portaudio
if ! brew list portaudio &>/dev/null; then
    info "Installing portaudio..."
    brew install portaudio
fi
ok "portaudio"

# Build
TMPDIR=$(mktemp -d)
trap "rm -rf $TMPDIR" EXIT

info "Downloading golos..."
git clone --depth 1 "https://github.com/$REPO.git" "$TMPDIR/golos" 2>/dev/null
ok "Downloaded"

info "Building..."
cd "$TMPDIR/golos"
go build -o golos .
ok "Built"

# Install binary
info "Installing to $INSTALL_DIR..."
if [[ -w "$INSTALL_DIR" ]]; then
    cp golos "$INSTALL_DIR/golos"
else
    sudo cp golos "$INSTALL_DIR/golos"
fi
chmod +x "$INSTALL_DIR/golos"
ok "Installed to $INSTALL_DIR/golos"

# Create config directory
mkdir -p "$CONFIG_DIR"

# Prompt for Deepgram API key if not set
if [[ -z "$DEEPGRAM_API_KEY" ]] && [[ ! -f "$CONFIG_DIR/config.toml" ]]; then
    echo ""
    echo -e "${YELLOW}Deepgram API key required for speech-to-text.${NC}"
    echo "Get one free at: https://console.deepgram.com"
    echo ""
    read -p "Enter your Deepgram API key (or press Enter to skip): " api_key
    if [[ -n "$api_key" ]]; then
        cat > "$CONFIG_DIR/config.toml" <<EOF
deepgram_api_key = "$api_key"
hotkey = "right_option"
output_mode = "clipboard"
language = "en-US"
EOF
        ok "Config saved to $CONFIG_DIR/config.toml"
    else
        warn "Skipped — set DEEPGRAM_API_KEY env var or create $CONFIG_DIR/config.toml later"
    fi
fi

echo ""
info "golos installed successfully!"
echo ""
echo "  Usage:"
echo "    golos        Run in foreground"
echo "    golos -d     Run in background (detached)"
echo "    golos stop   Stop background process"
echo ""
echo "  You may need to grant Accessibility permission to your terminal:"
echo "  System Settings → Privacy & Security → Accessibility"
echo ""
