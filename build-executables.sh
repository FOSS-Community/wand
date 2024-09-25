#!/bin/sh
# Install script for your Go CLI tool

# Define the CLI tool name and the GitHub repo where releases are hosted
CLI_NAME="wand"
REPO_URL="https://github.com/FOSS-Community/wand"

# Function to detect the OS and architecture
detect_os_arch() {
    OS=$(uname | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)

    # Determine architecture
    if [ "$ARCH" = "x86_64" ]; then
        ARCH="amd64"
    elif [ "$ARCH" = "arm64" ] || [ "$ARCH" = "aarch64" ]; then
        ARCH="arm64"
    elif [ "$ARCH" = "arm" ]; then
        ARCH="arm"
    else
        echo "Unsupported architecture: $ARCH"
        exit 1
    fi

    # Determine OS
    if [ "$OS" = "linux" ]; then
        OS="linux"
    elif [ "$OS" = "darwin" ]; then
        OS="macOS"
    else
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

    # Construct the correct download URL based on OS and architecture
    DOWNLOAD_URL="$REPO_URL/releases/download/$LATEST_VERSION/wand_v${LATEST_VERSION}_${OS_ARCH}.gz"
    INSTALL_PATH="/usr/local/bin/$CLI_NAME"

    echo "Downloading $CLI_NAME version $LATEST_VERSION for $OS_ARCH..."
    curl -L "$DOWNLOAD_URL" -o "$CLI_NAME.gz"

    if [ $? -ne 0 ]; then
        echo "Error downloading the binary."
        exit 1
    fi

    echo "Decompressing the binary..."
    gunzip "$CLI_NAME.gz"

    if [ $? -ne 0 ]; then
        echo "Error decompressing the binary."
        exit 1
    fi

    echo "Making the binary executable..."
    chmod +x "$CLI_NAME"

    echo "Moving the binary to $INSTALL_PATH..."
    sudo mv "$CLI_NAME" "$INSTALL_PATH"

    echo "$CLI_NAME has been installed successfully!"
}

install_cli
