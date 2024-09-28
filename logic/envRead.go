package logic

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// ReadEnvValue reads a value from a specified .env file
func ReadEnvValue(input interface{}, params ...interface{}) (interface{}, error) {
	if len(params) < 1 {
		return nil, fmt.Errorf("missing parameter: config file name")
	}

	key, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("expected string input for key")
	}

	fileName, ok := params[0].(string)
	if !ok {
		return nil, fmt.Errorf("expected string input for file name")
	}

	// Load the specified .env file
	err := godotenv.Load(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	// Get the value for the key
	value := os.Getenv(key)
	if value == "" {
		return nil, fmt.Errorf("config key %s not found in %s", key, fileName)
	}

	return value, nil
}
