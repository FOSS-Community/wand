#!/bin/sh
# Install script for your Go CLI tool

# Define the CLI tool name and the GitHub repo where releases are hosted
CLI_NAME="wand"
REPO_URL="https://github.com/FOSS-Community/wand"

# Function to detect the OS and architecture
detect_os_arch() {
    OS=$(uname | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)

    # Check for supported architectures
    if [ "$ARCH" = "x86_64" ]; then
        ARCH="amd64"
    elif [ "$ARCH" = "arm64" ] || [ "$ARCH" = "aarch64" ]; then
        ARCH="arm64"
    else
        echo "Unsupported architecture: $ARCH"
        exit 1
    fi

    # Check for supported operating systems
    if [ "$OS" != "linux" ] && [ "$OS" != "darwin" ]; then
        echo "Unsupported OS: $OS"
        exit 1
    fi

    echo "$OS-$ARCH"
}

# Function to download and install the binary
install_cli() {
    OS_ARCH=$(detect_os_arch)
    LATEST_VERSION=$(curl -s "$REPO_URL/releases/latest" | grep -oP 'tag/\Kv[0-9.]+')

    if [ -z "$LATEST_VERSION" ]; then
        echo "Unable to find the latest release version."
        exit 1
    fi

    DOWNLOAD_URL="$REPO_URL/releases/download/$LATEST_VERSION/$CLI_NAME-$OS_ARCH"
    INSTALL_PATH="/usr/local/bin/$CLI_NAME"

    echo "Downloading $CLI_NAME version $LATEST_VERSION for $OS_ARCH..."
    curl -L "$DOWNLOAD_URL" -o "$CLI_NAME"

    if [ $? -ne 0 ]; then
        echo "Error downloading the binary."
        exit 1
    fi

    echo "Making the binary executable..."
    chmod +x "$CLI_NAME"

    echo "Moving the binary to $INSTALL_PATH..."
    sudo mv "$CLI_NAME" "$INSTALL_PATH"

    echo "$CLI_NAME has been installed successfully!"
}

install_cli