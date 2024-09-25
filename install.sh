#!/bin/sh
# Install script for your Go CLI tool

# Define the CLI tool name, GitHub repo, and manual version number
CLI_NAME="wand"
REPO_URL="https://github.com/FOSS-Community/wand"
MANUAL_VERSION="v0.1.0"  # Replace this with the actual latest version

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
    
    DOWNLOAD_URL="$REPO_URL/releases/download/$MANUAL_VERSION/$CLI_NAME-$MANUAL_VERSION-$OS_ARCH.tar.gz"
    INSTALL_PATH="/usr/local/bin/$CLI_NAME"

    echo "Downloading $CLI_NAME version $MANUAL_VERSION for $OS_ARCH..."
    curl -L "$DOWNLOAD_URL" -o "$CLI_NAME.tar.gz"
    if [ $? -ne 0 ]; then
        echo "Error downloading the archive."
        exit 1
    fi

    echo "Extracting the binary..."
    tar -xzf "$CLI_NAME.tar.gz"
    if [ $? -ne 0 ]; then
        echo "Error extracting the archive."
        exit 1
    fi

    echo "Making the binary executable..."
    chmod +x "$CLI_NAME"

    echo "Moving the binary to $INSTALL_PATH..."
    sudo mv "$CLI_NAME" "$INSTALL_PATH"

    echo "Cleaning up..."
    rm "$CLI_NAME.tar.gz"

    echo "$CLI_NAME has been installed successfully!"
}

install_cli
