#!/bin/sh
set -e

# Determine OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64)
        ARCH="arm64"
        ;;
    i386|i686)
        ARCH="386"
        ;;
esac

# Set version
VERSION="0.0.1"

# Construct download URL
DOWNLOAD_URL="https://github.com/FOSS-Community/wand/releases/download/v0.0.1/wand_${VERSION}_${OS}_${ARCH}.tar.gz"

# Create temporary directory
TMP_DIR=$(mktemp -d)
trap 'rm -rf "$TMP_DIR"' EXIT

# Download and extract
echo "Downloading Wand ${VERSION} for ${OS} ${ARCH}..."
curl -L "$DOWNLOAD_URL" | tar -xz -C "$TMP_DIR"

# Install
INSTALL_DIR="/usr/local/bin"
echo "Installing Wand to $INSTALL_DIR..."
sudo mv "$TMP_DIR/wand" "$INSTALL_DIR/"

echo "Wand ${VERSION} has been installed successfully!"
echo "You can now use the 'wand' command."
