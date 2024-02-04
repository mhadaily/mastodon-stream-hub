#!/bin/bash

# Function to install protoc-gen-go and protoc-gen-go-grpc
install_go_protoc_plugins() {
    echo "Installing protoc-gen-go and protoc-gen-go-grpc..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
}

# Function to add $GOPATH/bin to PATH in the shell profile
update_path() {
    PROFILE_FILE=~/.bashrc
    if [[ "$SHELL" == *"zsh"* ]]; then
        PROFILE_FILE=~/.zshrc
    fi
    
    if ! grep -q 'export PATH=$PATH:$(go env GOPATH)/bin' "$PROFILE_FILE"; then
        echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> "$PROFILE_FILE"
        echo "Updated $PROFILE_FILE to include Go bin directory in PATH."
    fi
}

# Install protoc compiler if not present
if ! command -v protoc &> /dev/null; then
    echo "protoc could not be found, installing..."
    case "$OSTYPE" in
        linux*)
            sudo apt-get update && sudo apt-get install -y protobuf-compiler
            ;;
        darwin*)
            brew install protobuf
            ;;
        *)
            echo "Unsupported OS. Please manually install protoc."
            exit 1
            ;;
    esac
else
    echo "protoc is already installed."
fi

# Install Go plugins for protoc
install_go_protoc_plugins

# Update PATH to include $GOPATH/bin
update_path

# Inform the user to reload their shell
echo "Please reload your shell or run 'source ~/.bashrc' (or ~/.zshrc) to update your PATH."

