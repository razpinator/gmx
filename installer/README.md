# GMX Go-based Installer

A native Go installer for the GMX CLI tool that provides a cross-platform installation solution.

## Features

- **Smart Installation**: Uses `go install` when Go is available, falls back to binary downloads
- **Cross-platform**: Works on Linux, macOS, and Windows
- **Automatic PATH Configuration**: Guides users through PATH setup
- **Release Integration**: Can download from GitHub releases when Go is not available
- **Single Binary**: Self-contained installer that can be distributed independently

## Usage

### Build the Installer

```bash
cd installer
go build -o gmx-installer .
```

### Run the Installer

```bash
./gmx-installer
```

### Uninstall GMX

```bash
./gmx-installer uninstall
```

## How It Works

1. **Go Detection**: Checks if Go is installed on the system
2. **Installation Method**:
   - If Go is available: Uses `go install github.com/razpinator/gmx@latest`
   - If Go is not available: Downloads appropriate binary from GitHub releases
3. **PATH Configuration**: Automatically configures PATH and provides guidance
4. **Verification**: Tests the installation to ensure gmx works correctly

## Uninstallation

The installer also includes uninstall functionality:

1. **Automatic Removal**: Finds and removes gmx from common installation locations
2. **Cache Cleanup**: Cleans Go module cache entries
3. **PATH Cleanup**: Provides guidance for removing PATH configurations
4. **Verification**: Confirms gmx has been successfully removed

## Installation Locations

- **With Go**: Installs to `$(go env GOPATH)/bin`
- **Without Go**: Installs to `~/.local/bin` (Linux/macOS) or user directory (Windows)

## Building for Distribution

To create installers for different platforms:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o gmx-installer-linux-amd64 .

# macOS
GOOS=darwin GOARCH=amd64 go build -o gmx-installer-darwin-amd64 .
GOOS=darwin GOARCH=arm64 go build -o gmx-installer-darwin-arm64 .

# Windows
GOOS=windows GOARCH=amd64 go build -o gmx-installer-windows-amd64.exe .
```

This installer provides a native Go alternative to shell scripts for users who prefer Go-based tooling.