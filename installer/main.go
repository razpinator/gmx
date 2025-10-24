package main

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	repoOwner = "razpinator"
	repoName  = "gmx"
	apiURL    = "https://api.github.com/repos/" + repoOwner + "/" + repoName + "/releases/latest"
)

type Release struct {
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func main() {
	fmt.Println("🚀 GMX Go-based Installer")
	fmt.Println("========================")

	// Check if Go is available for fallback
	goAvailable := isGoAvailable()
	if goAvailable {
		fmt.Println("✅ Go detected - using go install method")
		installWithGo()
	} else {
		fmt.Println("⚠️  Go not detected - downloading binary release")
		installFromRelease()
	}

	// Test installation
	fmt.Println("\n🧪 Testing installation...")
	if testInstallation() {
		fmt.Println("✅ GMX installed successfully!")
		printUsage()
	} else {
		fmt.Println("❌ Installation failed")
		os.Exit(1)
	}
}

func isGoAvailable() bool {
	_, err := exec.LookPath("go")
	return err == nil
}

func installWithGo() {
	fmt.Println("📦 Installing gmx using 'go install'...")

	// Set up environment for better temp directory handling
	homeDir, _ := os.UserHomeDir()
	tempDir := filepath.Join(homeDir, ".cache", "go-build")
	os.MkdirAll(tempDir, 0755)
	os.Setenv("GOTMPDIR", tempDir)

	cmd := exec.Command("go", "install", "github.com/razpinator/gmx@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ go install failed: %v\n", err)
		fmt.Println("🔄 Falling back to binary release...")
		installFromRelease()
		return
	}

	// Configure PATH
	configurePath()
}

func installFromRelease() {
	fmt.Println("📥 Fetching latest release information...")

	release, err := getLatestRelease()
	if err != nil {
		fmt.Printf("❌ Failed to fetch release info: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("📋 Latest version: %s\n", release.TagName)

	// Find appropriate asset for current platform
	assetName := getBinaryName()
	var downloadURL string

	for _, asset := range release.Assets {
		if strings.Contains(asset.Name, assetName) {
			downloadURL = asset.BrowserDownloadURL
			break
		}
	}

	if downloadURL == "" {
		fmt.Printf("❌ No binary found for %s-%s\n", runtime.GOOS, runtime.GOARCH)
		os.Exit(1)
	}

	fmt.Printf("⬇️  Downloading %s...\n", assetName)
	if err := downloadAndInstall(downloadURL); err != nil {
		fmt.Printf("❌ Download failed: %v\n", err)
		os.Exit(1)
	}

	configurePath()
}

func getLatestRelease() (*Release, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	return &release, nil
}

func getBinaryName() string {
	os := runtime.GOOS
	arch := runtime.GOARCH

	// Convert Go arch names to common naming
	switch arch {
	case "amd64":
		arch = "x86_64"
	case "arm64":
		arch = "aarch64"
	}

	return fmt.Sprintf("%s-%s", os, arch)
}

func downloadAndInstall(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create installation directory
	installDir := getInstallDir()
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return err
	}

	// Determine if it's a compressed archive or direct binary
	if strings.HasSuffix(url, ".tar.gz") {
		return extractTarGz(resp.Body, installDir)
	} else {
		// Direct binary download
		binaryName := "gmx"
		if runtime.GOOS == "windows" {
			binaryName += ".exe"
		}

		binaryPath := filepath.Join(installDir, binaryName)
		file, err := os.Create(binaryPath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return err
		}

		// Make executable on Unix-like systems
		if runtime.GOOS != "windows" {
			err = os.Chmod(binaryPath, 0755)
		}

		return err
	}
}

func extractTarGz(src io.Reader, destDir string) error {
	gzr, err := gzip.NewReader(src)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if header.Name == "gmx" || header.Name == "gmx.exe" {
			path := filepath.Join(destDir, header.Name)
			file, err := os.Create(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(file, tr)
			if err != nil {
				return err
			}

			// Make executable
			if runtime.GOOS != "windows" {
				err = os.Chmod(path, 0755)
			}

			return err
		}
	}

	return fmt.Errorf("gmx binary not found in archive")
}

func getInstallDir() string {
	if isGoAvailable() {
		// Use Go's bin directory if available
		cmd := exec.Command("go", "env", "GOPATH")
		output, err := cmd.Output()
		if err == nil {
			gopath := strings.TrimSpace(string(output))
			return filepath.Join(gopath, "bin")
		}
	}

	// Fallback to user's local bin
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".local", "bin")
}

func configurePath() {
	installDir := getInstallDir()
	fmt.Printf("📁 Binary installed to: %s\n", installDir)

	// Add to PATH for current session
	currentPath := os.Getenv("PATH")
	if !strings.Contains(currentPath, installDir) {
		newPath := installDir + string(os.PathListSeparator) + currentPath
		os.Setenv("PATH", newPath)
		fmt.Println("✅ Added to PATH for current session")
	}

	// Suggest permanent PATH addition
	fmt.Println("\n💡 To permanently add gmx to your PATH:")

	switch runtime.GOOS {
	case "windows":
		fmt.Printf("   Add '%s' to your system PATH\n", installDir)
	case "darwin", "linux":
		shell := os.Getenv("SHELL")
		if strings.Contains(shell, "zsh") {
			fmt.Printf("   echo 'export PATH=\"%s:$PATH\"' >> ~/.zshrc\n", installDir)
		} else {
			fmt.Printf("   echo 'export PATH=\"%s:$PATH\"' >> ~/.bashrc\n", installDir)
		}
	}
}

func testInstallation() bool {
	binaryName := "gmx"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	// Try to find gmx in PATH
	_, err := exec.LookPath(binaryName)
	if err == nil {
		return true
	}

	// Try in install directory
	installDir := getInstallDir()
	binaryPath := filepath.Join(installDir, binaryName)
	if _, err := os.Stat(binaryPath); err == nil {
		return true
	}

	return false
}

func printUsage() {
	fmt.Println("\n📚 Usage:")
	fmt.Println("  gmx init           - Initialize a new project")
	fmt.Println("  gmx run <workflow> - Run a workflow")
	fmt.Println("  gmx --help         - Show help")
	fmt.Println("\n🌟 Happy coding with GMX!")
}
