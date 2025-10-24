package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/razpinator/gmx/models"
	"gopkg.in/yaml.v2"
)

// Config holds the filenames for data, template, and output
type Config struct {
	Items []models.Item `yaml:"items"`
}

// readJSON reads and parses the JSON data file
func ReadJSON(filename string) map[string]interface{} {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read data file: %v", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalf("Failed to parse data file: %v", err)
	}

	return data
}

// readFile reads the content of a file
func ReadFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read template file: %v", err)
	}

	return string(content)
}

func WriteFileWithCustomSeparator(filePath string, data []byte, perm os.FileMode) error {
	// Replace custom path separator with OS-specific path separator
	normalizedPath := strings.ReplaceAll(filePath, ">", string(os.PathSeparator))

	// Get the directory path
	dirPath := filepath.Dir(normalizedPath)

	// Create directories if they don't exist
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	// Write the file
	if err := os.WriteFile(normalizedPath, data, perm); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// readConfig reads and parses the YAML configuration file
func ReadConfig(filename string) Config {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	return config
}
