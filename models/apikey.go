package models

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateAPIKey returns a fresh API key for a user, drawn from the system CSPRNG.
func GenerateAPIKey() (string, error) {
	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}
