# GMX Uninstaller Script for Windows PowerShell
# This script removes gmx and cleans up PATH configuration

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

Write-ColorText "GMX Uninstaller for Windows" "Cyan"
Write-ColorText "============================" "Cyan"
Write-Host ""

$removed = $false

# Function to remove Go bin directory from PATH
function Remove-GoFromPath {
    Write-ColorText "Cleaning PATH configuration..." "Cyan"
    
    $currentPath = [Environment]::GetEnvironmentVariable("PATH", [EnvironmentVariableTarget]::User)
    
    if ($currentPath) {
        # Get Go bin directory
        try {
            $goPath = go env GOPATH
            $goBinDir = Join-Path $goPath "bin"
            
            # Remove Go bin directory from PATH
            $pathArray = $currentPath -split ";"
            $newPathArray = $pathArray | Where-Object { $_ -ne $goBinDir }
            
            if ($pathArray.Count -ne $newPathArray.Count) {
                $newPath = $newPathArray -join ";"
                [Environment]::SetEnvironmentVariable("PATH", $newPath, [EnvironmentVariableTarget]::User)
                Write-ColorText "✓ Removed Go bin directory from user PATH" "Green"
                
                # Update PATH for current session
                $env:PATH = $env:PATH.Replace(";$goBinDir", "").Replace("$goBinDir;", "")
            } else {
                Write-ColorText "! Go bin directory not found in user PATH" "Yellow"
            }
        } catch {
            Write-ColorText "! Could not get Go path: $_" "Yellow"
        }
    }
}

# Find and remove gmx binary
Write-ColorText "Looking for gmx installations..." "Cyan"

$locations = @()

# Add Go bin directory if available
try {
    $goPath = go env GOPATH
    $goBinDir = Join-Path $goPath "bin"
    $locations += $goBinDir
} catch {
    Write-ColorText "! Go not available" "Yellow"
}

# Add other common locations
$userProfile = $env:USERPROFILE
$locations += @(
    (Join-Path $userProfile ".local\bin"),
    (Join-Path $userProfile "bin"),
    "C:\Program Files\gmx",
    "C:\gmx"
)

# Try to remove from each location
foreach ($location in $locations) {
    $gmxPath = Join-Path $location "gmx.exe"
    if (Test-Path $gmxPath) {
        Write-ColorText "Found gmx at: $gmxPath" "Yellow"
        try {
            Remove-Item $gmxPath -Force
            Write-ColorText "✓ Removed gmx from $location" "Green"
            $removed = $true
        } catch {
            Write-ColorText "✗ Failed to remove $gmxPath : $_" "Red"
            Write-ColorText "! Try running as Administrator" "Yellow"
        }
    }
}

# Clean Go module cache if available
try {
    Write-Host ""
    Write-ColorText "Cleaning Go module cache..." "Cyan"
    go clean -modcache github.com/razpinator/gmx
    Write-ColorText "✓ Cleaned Go module cache" "Green"
} catch {
    Write-ColorText "! Could not clean Go module cache" "Yellow"
}

# Remove from PATH
Remove-GoFromPath

# Check if gmx is still accessible
Write-Host ""
Write-ColorText "Verifying removal..." "Cyan"

try {
    $gmxVersion = gmx --version 2>$null
    Write-ColorText "! gmx is still accessible - you may need to restart PowerShell" "Yellow"
    Write-ColorText "! Or there may be another installation in a different location" "Yellow"
} catch {
    Write-ColorText "✓ gmx is no longer accessible" "Green"
    $removed = $true
}

# Final message
if ($removed) {
    Write-Host ""
    Write-ColorText "✅ GMX has been successfully uninstalled!" "Green"
    Write-Host ""
    Write-ColorText "Note: Restart your PowerShell/Command Prompt to ensure PATH changes take effect" "Yellow"
    Write-Host ""
    Write-ColorText "👋 Thanks for using GMX!" "Cyan"
} else {
    Write-Host ""
    Write-ColorText "❌ No gmx installation found to remove" "Red"
    Write-ColorText "gmx may have been installed manually or to a custom location" "Yellow"
    Write-ColorText "You can manually check with: where gmx" "Yellow"
}