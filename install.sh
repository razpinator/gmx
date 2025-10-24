#!/bin/bash

# GMX Installation Script
# This script installs gmx and automatically configures PATH

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Convert architecture names
case $ARCH in
    x86_64) ARCH="amd64" ;;
    arm64|aarch64) ARCH="arm64" ;;
    i386) ARCH="386" ;;
    *) echo -e "${RED}Unsupported architecture: $ARCH${NC}"; exit 1 ;;
esac

echo -e "${BLUE}GMX Installer${NC}"
echo -e "${BLUE}==============${NC}"
echo -e "Detected OS: ${YELLOW}$OS${NC}"
echo -e "Detected Architecture: ${YELLOW}$ARCH${NC}"
echo ""

# Check if Go is installed
if command -v go &> /dev/null; then
    echo -e "${GREEN}✓${NC} Go is installed"
    GO_VERSION=$(go version | cut -d' ' -f3)
    echo -e "  Go version: ${YELLOW}$GO_VERSION${NC}"
    
    # Install using go install (preferred method)
    echo -e "\n${BLUE}Installing gmx using 'go install'...${NC}"
    go install github.com/razpinator/gmx@latest
    
    # Get Go paths
    GOPATH=$(go env GOPATH)
    GOBIN_DIR="$GOPATH/bin"
    
    # Check if binary was installed
    if [[ -f "$GOBIN_DIR/gmx" ]]; then
        echo -e "${GREEN}✓${NC} gmx installed successfully to $GOBIN_DIR/gmx"
    else
        echo -e "${RED}✗${NC} Installation failed"
        exit 1
    fi
    
else
    echo -e "${YELLOW}!${NC} Go is not installed"
    echo -e "${RED}Please install Go first: https://golang.org/dl/${NC}"
    exit 1
fi

# Function to add PATH to shell config
add_to_path() {
    local shell_config="$1"
    local path_export="export PATH=\$PATH:\$(go env GOPATH)/bin"
    
    if [[ -f "$shell_config" ]]; then
        # Check if PATH is already configured
        if grep -q "go env GOPATH" "$shell_config" 2>/dev/null; then
            echo -e "${YELLOW}!${NC} PATH already configured in $shell_config"
        else
            echo -e "\n# Added by GMX installer" >> "$shell_config"
            echo "$path_export" >> "$shell_config"
            echo -e "${GREEN}✓${NC} Added Go bin directory to PATH in $shell_config"
        fi
    else
        # Create the file if it doesn't exist
        echo "# Added by GMX installer" > "$shell_config"
        echo "$path_export" >> "$shell_config"
        echo -e "${GREEN}✓${NC} Created $shell_config and added Go bin directory to PATH"
    fi
}

# Configure PATH for different shells
echo -e "\n${BLUE}Configuring PATH...${NC}"

# Detect current shell
CURRENT_SHELL=$(basename "$SHELL")
echo -e "Current shell: ${YELLOW}$CURRENT_SHELL${NC}"

case $CURRENT_SHELL in
    bash)
        add_to_path "$HOME/.bashrc"
        add_to_path "$HOME/.bash_profile"
        ;;
    zsh)
        add_to_path "$HOME/.zshrc"
        ;;
    fish)
        # Fish shell has different syntax
        FISH_CONFIG="$HOME/.config/fish/config.fish"
        if [[ -f "$FISH_CONFIG" ]]; then
            if grep -q "go env GOPATH" "$FISH_CONFIG" 2>/dev/null; then
                echo -e "${YELLOW}!${NC} PATH already configured in $FISH_CONFIG"
            else
                echo -e "\n# Added by GMX installer" >> "$FISH_CONFIG"
                echo "set -gx PATH \$PATH (go env GOPATH)/bin" >> "$FISH_CONFIG"
                echo -e "${GREEN}✓${NC} Added Go bin directory to PATH in $FISH_CONFIG"
            fi
        else
            mkdir -p "$(dirname "$FISH_CONFIG")"
            echo "# Added by GMX installer" > "$FISH_CONFIG"
            echo "set -gx PATH \$PATH (go env GOPATH)/bin" >> "$FISH_CONFIG"
            echo -e "${GREEN}✓${NC} Created $FISH_CONFIG and added Go bin directory to PATH"
        fi
        ;;
    *)
        echo -e "${YELLOW}!${NC} Unknown shell: $CURRENT_SHELL"
        echo -e "${YELLOW}!${NC} Please manually add \$(go env GOPATH)/bin to your PATH"
        ;;
esac

# Test installation
echo -e "\n${BLUE}Testing installation...${NC}"

# Add Go bin to current session PATH
export PATH="$PATH:$(go env GOPATH)/bin"

if command -v gmx &> /dev/null; then
    echo -e "${GREEN}✓${NC} gmx is working correctly!"
    echo -e "\n${GREEN}Installation completed successfully!${NC}"
    echo -e "\n${BLUE}Usage:${NC}"
    echo -e "  ${YELLOW}gmx init${NC}        - Initialize a new project"
    echo -e "  ${YELLOW}gmx run <workflow>${NC} - Run a workflow"
    echo -e "  ${YELLOW}gmx --help${NC}      - Show help"
    echo -e "\n${YELLOW}Note:${NC} Restart your terminal or run 'source ~/.zshrc' (or your shell config) to use gmx in new sessions."
else
    echo -e "${RED}✗${NC} gmx command not found. Installation may have failed."
    echo -e "${YELLOW}Try running:${NC} source ~/.zshrc (or your shell config file)"
    exit 1
fi