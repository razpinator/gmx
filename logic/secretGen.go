package logic

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GenerateSecret generates a 16-byte hexadecimal secret
func Generate16bitSecret(input interface{}, params ...interface{}) (interface{}, error) {
	secret := make([]byte, 16)
	_, err := rand.Read(secret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate secret: %w", err)
	}
	return hex.EncodeToString(secret), nil
}

// Generate64BitSecret generates an 8-byte (64-bit) hexadecimal secret
func Generate64BitSecret(input interface{}, params ...interface{}) (interface{}, error) {
	secret := make([]byte, 8)
	_, err := rand.Read(secret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate secret: %w", err)
	}
	return hex.EncodeToString(secret), nil
}
