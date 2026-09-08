package models

import (
	"math/rand"
)

const apiKeyCharset = "abcdef0123456789"

// GenerateAPIKey returns a fresh API key for a user.
func GenerateAPIKey() string {
	key := make([]byte, 32)
	for i := range key {
		key[i] = apiKeyCharset[rand.Intn(len(apiKeyCharset))]
	}
	return string(key)
}
