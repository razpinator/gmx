# GMX Installation Script for Windows PowerShell
# This script installs gmx and automatically configures PATH

param(
    [switch]$Force
)

# Set error action preference
$ErrorActionPreference = "Stop"

# Colors for output (PowerShell 5.0+)
function Write-ColorText {
    param(
        [string]$Text,
        [string]$Color = "White"
    )
    if ($PSVersionTable.PSVersion.Major -ge 5) {
        Write-Host $Text -ForegroundColor $Color
    } else {
        Write-Host $Text
    }
}

Write-ColorText "GMX Installer for Windows" "Cyan"
Write-ColorText "==========================" "Cyan"
Write-Host ""

# Check if Go is installed
try {
    $goVersion = go version
    Write-ColorText "✓ Go is installed" "Green"
    Write-ColorText "  Go version: $goVersion" "Yellow"
} catch {
    Write-ColorText "✗ Go is not installed" "Red"
    Write-ColorText "Please install Go first: https://golang.org/dl/" "Red"
    exit 1
}

# Install using go install
Write-Host ""
Write-ColorText "Installing gmx using 'go install'..." "Cyan"

try {
    go install github.com/razpinator/gmx@latest
    Write-ColorText "✓ gmx installed successfully" "Green"
} catch {
    Write-ColorText "✗ Installation failed: $_" "Red"
    exit 1
}

# Get Go paths
$goPath = go env GOPATH
$goBinDir = Join-Path $goPath "bin"
$gmxPath = Join-Path $goBinDir "gmx.exe"

# Check if binary was installed
if (Test-Path $gmxPath) {
    Write-ColorText "✓ gmx binary found at $gmxPath" "Green"
} else {
    Write-ColorText "✗ gmx binary not found" "Red"
    exit 1
}

# Configure PATH
Write-Host ""
Write-ColorText "Configuring PATH..." "Cyan"

# Get current user PATH
$currentPath = [Environment]::GetEnvironmentVariable("PATH", [EnvironmentVariableTarget]::User)

# Check if Go bin directory is already in PATH
if ($currentPath -split ";" -contains $goBinDir) {
    Write-ColorText "! Go bin directory already in PATH" "Yellow"
} else {
    # Add Go bin directory to user PATH
    $newPath = if ($currentPath) { "$currentPath;$goBinDir" } else { $goBinDir }
    [Environment]::SetEnvironmentVariable("PATH", $newPath, [EnvironmentVariableTarget]::User)
    Write-ColorText "✓ Added Go bin directory to user PATH" "Green"
    
    # Update PATH for current session
    $env:PATH = "$env:PATH;$goBinDir"
}

# Test installation
Write-Host ""
Write-ColorText "Testing installation..." "Cyan"

try {
    $testResult = & $gmxPath --help
    if ($LASTEXITCODE -eq 0) {
        Write-ColorText "✓ gmx is working correctly!" "Green"
        Write-Host ""
        Write-ColorText "Installation completed successfully!" "Green"
        Write-Host ""
        Write-ColorText "Usage:" "Cyan"
        Write-ColorText "  gmx init        - Initialize a new project" "Yellow"
        Write-ColorText "  gmx run <workflow> - Run a workflow" "Yellow"
        Write-ColorText "  gmx --help      - Show help" "Yellow"
        Write-Host ""
        Write-ColorText "Note: You may need to restart your terminal/PowerShell to use gmx in new sessions." "Yellow"
    } else {
        throw "gmx returned non-zero exit code"
    }
} catch {
    Write-ColorText "✗ gmx test failed: $_" "Red"
    Write-ColorText "Try restarting your terminal/PowerShell" "Yellow"
    exit 1
}