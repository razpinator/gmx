#!/bin/bash

# GMX Uninstaller Script
# This script removes gmx and cleans up PATH configuration

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}GMX Uninstaller${NC}"
echo -e "${BLUE}===============${NC}"
echo ""

removed=false

# Function to remove PATH entries from shell config
remove_from_path() {
    local shell_config="$1"
    
    if [[ -f "$shell_config" ]]; then
        # Check if GMX installer added PATH configuration
        if grep -q "# Added by GMX installer" "$shell_config" 2>/dev/null; then
            echo -e "${YELLOW}Removing PATH configuration from $shell_config...${NC}"
            
            # Create a backup
            cp "$shell_config" "$shell_config.backup.$(date +%Y%m%d_%H%M%S)"
            
            # Remove GMX installer lines
            sed -i.tmp '/# Added by GMX installer/,+1d' "$shell_config" 2>/dev/null || true
            rm -f "$shell_config.tmp" 2>/dev/null || true
            
            echo -e "${GREEN}✓${NC} Removed PATH configuration from $shell_config"
            echo -e "${BLUE}Backup created: $shell_config.backup.*${NC}"
        else
            echo -e "${YELLOW}!${NC} No GMX installer PATH configuration found in $shell_config"
        fi
    fi
}

# Find and remove gmx binary
echo -e "${BLUE}Looking for gmx installations...${NC}"

# Check Go bin directory first
if command -v go &> /dev/null; then
    GOPATH=$(go env GOPATH)
    GOBIN_DIR="$GOPATH/bin"
    
    if [[ -f "$GOBIN_DIR/gmx" ]]; then
        echo -e "${YELLOW}Found gmx in Go bin directory: $GOBIN_DIR/gmx${NC}"
        rm "$GOBIN_DIR/gmx"
        echo -e "${GREEN}✓${NC} Removed gmx from $GOBIN_DIR"
        removed=true
    fi
fi

# Check other common locations
locations=(
    "$HOME/.local/bin"
    "$HOME/bin"
    "/usr/local/bin"
    "/usr/bin"
)

for location in "${locations[@]}"; do
    if [[ -f "$location/gmx" ]]; then
        echo -e "${YELLOW}Found gmx in: $location/gmx${NC}"
        if rm "$location/gmx" 2>/dev/null; then
            echo -e "${GREEN}✓${NC} Removed gmx from $location"
            removed=true
        else
            echo -e "${RED}✗${NC} Failed to remove $location/gmx (permission denied)"
            echo -e "${YELLOW}Try: sudo rm $location/gmx${NC}"
        fi
    fi
done

# Clean Go module cache if Go is available
if command -v go &> /dev/null; then
    echo -e "\n${BLUE}Cleaning Go module cache...${NC}"
    go clean -modcache github.com/razpinator/gmx 2>/dev/null || true
    echo -e "${GREEN}✓${NC} Cleaned Go module cache"
fi

# Remove PATH configuration
echo -e "\n${BLUE}Cleaning PATH configuration...${NC}"

# Detect current shell and clean config
CURRENT_SHELL=$(basename "$SHELL")

case $CURRENT_SHELL in
    bash)
        remove_from_path "$HOME/.bashrc"
        remove_from_path "$HOME/.bash_profile"
        ;;
    zsh)
        remove_from_path "$HOME/.zshrc"
        ;;
    fish)
        FISH_CONFIG="$HOME/.config/fish/config.fish"
        if [[ -f "$FISH_CONFIG" ]] && grep -q "# Added by GMX installer" "$FISH_CONFIG" 2>/dev/null; then
            echo -e "${YELLOW}Removing PATH configuration from $FISH_CONFIG...${NC}"
            cp "$FISH_CONFIG" "$FISH_CONFIG.backup.$(date +%Y%m%d_%H%M%S)"
            sed -i.tmp '/# Added by GMX installer/,+1d' "$FISH_CONFIG" 2>/dev/null || true
            rm -f "$FISH_CONFIG.tmp" 2>/dev/null || true
            echo -e "${GREEN}✓${NC} Removed PATH configuration from $FISH_CONFIG"
        fi
        ;;
    *)
        echo -e "${YELLOW}!${NC} Unknown shell: $CURRENT_SHELL"
        echo -e "${YELLOW}!${NC} Please manually remove Go bin PATH entries from your shell config"
        ;;
esac

# Check if gmx is still accessible
echo -e "\n${BLUE}Verifying removal...${NC}"
if command -v gmx &> /dev/null; then
    echo -e "${YELLOW}!${NC} gmx is still in PATH - you may need to restart your terminal"
    echo -e "${YELLOW}!${NC} Or there may be another installation in a different location"
else
    echo -e "${GREEN}✓${NC} gmx is no longer accessible"
    removed=true
fi

# Final message
if $removed; then
    echo -e "\n${GREEN}✅ GMX has been successfully uninstalled!${NC}"
    echo -e "\n${YELLOW}Note:${NC} Restart your terminal to ensure PATH changes take effect"
    echo -e "${BLUE}Shell config backups were created with timestamp suffixes${NC}"
    echo -e "\n${BLUE}👋 Thanks for using GMX!${NC}"
else
    echo -e "\n${RED}❌ No gmx installation found to remove${NC}"
    echo -e "${YELLOW}gmx may have been installed manually or to a custom location${NC}"
    echo -e "${YELLOW}You can manually check: which gmx${NC}"
fi