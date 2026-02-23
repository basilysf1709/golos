#!/bin/bash
set -e

REPO="basilysf1709/golos"
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

# Install portaudio (required at runtime)
if ! brew list portaudio &>/dev/null; then
    info "Installing portaudio..."
    brew install portaudio
fi
ok "portaudio"

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
    x86_64) GOARCH="amd64" ;;
    arm64)  GOARCH="arm64" ;;
    *)      fail "Unsupported architecture: $ARCH" ;;
esac

# Try downloading pre-built binary from GitHub Releases
TMPDIR=$(mktemp -d)
trap "rm -rf $TMPDIR" EXIT

LATEST_TAG=$(curl -sL "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name"' | head -1 | cut -d'"' -f4)

if [[ -n "$LATEST_TAG" ]]; then
    VERSION="${LATEST_TAG#v}"
    TARBALL="golos_${VERSION}_darwin_${GOARCH}.tar.gz"
    URL="https://github.com/$REPO/releases/download/${LATEST_TAG}/${TARBALL}"

    info "Downloading golos $LATEST_TAG..."
    if curl -sL --fail -o "$TMPDIR/$TARBALL" "$URL" 2>/dev/null; then
        tar -xzf "$TMPDIR/$TARBALL" -C "$TMPDIR"
        ok "Downloaded $LATEST_TAG"
    else
        warn "Pre-built binary not available, building from source..."
        LATEST_TAG=""
    fi
fi

# Fallback: build from source
if [[ -z "$LATEST_TAG" ]]; then
    # Check Go
    if ! command -v go &>/dev/null; then
        info "Installing Go..."
        brew install go
    fi
    ok "Go $(go version | awk '{print $3}' | sed 's/go//')"

    info "Downloading source..."
    git clone --depth 1 "https://github.com/$REPO.git" "$TMPDIR/golos" 2>/dev/null
    ok "Downloaded"

    info "Building..."
    cd "$TMPDIR/golos"
    go build -o "$TMPDIR/golos_bin" .
    mv "$TMPDIR/golos_bin" "$TMPDIR/golos_binary"
    # Clean up cloned dir so the binary path is clear
    rm -rf "$TMPDIR/golos"
    mv "$TMPDIR/golos_binary" "$TMPDIR/golos"
    ok "Built"
fi

# Install binary
info "Installing to $INSTALL_DIR..."
if [[ -w "$INSTALL_DIR" ]]; then
    cp "$TMPDIR/golos" "$INSTALL_DIR/golos"
else
    sudo cp "$TMPDIR/golos" "$INSTALL_DIR/golos"
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
