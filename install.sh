#!/bin/bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Print with color
print() {
    echo -e "${2}${1}${NC}"
}

# Error handler
handle_error() {
    print "Error: $1" "$RED"
    exit 1
}

# Detect OS and architecture
detect_platform() {
    OS="$(uname -s)"
    ARCH="$(uname -m)"
    
    case "$OS" in
        Linux*)     PLATFORM="linux" ;;
        Darwin*)    PLATFORM="darwin" ;;
        *)          handle_error "Unsupported operating system: $OS" ;;
    esac
    
    case "$ARCH" in
        x86_64)     ARCH="amd64" ;;
        arm64)      ARCH="arm64" ;;
        *)          handle_error "Unsupported architecture: $ARCH" ;;
    esac
    
    echo "${PLATFORM}-${ARCH}"
}

# Get latest version
get_latest_version() {
    curl -s https://api.github.com/repos/carakacloud/ksema-cli/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'
}

# Main installation process
main() {
    print "Installing ksema-cli..." "$YELLOW"
    
    # Detect platform
    PLATFORM=$(detect_platform)
    print "Detected platform: $PLATFORM" "$GREEN"
    
    # Get latest version
    VERSION=$(get_latest_version)
    print "Latest version: $VERSION" "$GREEN"
    
    # Create temp directory
    TEMP_DIR=$(mktemp -d)
    trap 'rm -rf "$TEMP_DIR"' EXIT
    
    # Download binary
    print "Downloading ksema-cli..." "$YELLOW"
    BINARY_NAME="ksema-cli-${PLATFORM}"
    if [ "$PLATFORM" = "windows-amd64" ]; then
        BINARY_NAME="${BINARY_NAME}.exe"
    fi
    
    DOWNLOAD_URL="https://github.com/carakacloud/ksema-cli/releases/download/${VERSION}/${BINARY_NAME}"
    
    if ! curl -L "$DOWNLOAD_URL" -o "$TEMP_DIR/ksema-cli"; then
        handle_error "Failed to download ksema-cli"
    fi
    
    # Make binary executable
    chmod +x "$TEMP_DIR/ksema-cli"
    
    # Install binary
    print "Installing to /usr/local/bin..." "$YELLOW"
    if ! sudo mv "$TEMP_DIR/ksema-cli" "/usr/local/bin/ksema-cli"; then
        handle_error "Failed to install ksema-cli"
    fi
    
    print "Installation complete! You can now use 'ksema-cli'" "$GREEN"
}

# Run main function
main 