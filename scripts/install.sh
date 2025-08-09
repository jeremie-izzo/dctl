#!/bin/bash

set -e

CLI_NAME="godev"
INSTALL_DIR="$HOME/bin"

echo "🔧 Building $CLI_NAME..."
go build -a -o $CLI_NAME

echo "📁 Ensuring $INSTALL_DIR exists..."
mkdir -p "$INSTALL_DIR"

echo "🚚 Moving $CLI_NAME to $INSTALL_DIR..."
mv "$CLI_NAME" "$INSTALL_DIR/"

# Ensure ~/bin is in PATH
if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
  echo "🔧 Adding $INSTALL_DIR to PATH in shell config..."

  SHELL_RC=""
  if [ -n "$ZSH_VERSION" ]; then
    SHELL_RC="$HOME/.zshrc"
  elif [ -n "$BASH_VERSION" ]; then
    SHELL_RC="$HOME/.bashrc"
  else
    SHELL_RC="$HOME/.profile"
  fi

  echo "export PATH=\"\$HOME/bin:\$PATH\"" >> "$SHELL_RC"
  echo "✅ Added to $SHELL_RC. Please restart your shell or run:"
  echo "source $SHELL_RC"
else
  echo "✅ $INSTALL_DIR already in PATH."
fi

echo "🎉 Installed! You can now run '$CLI_NAME' from anywhere."