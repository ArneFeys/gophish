package models

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// ValidateWebhookSignature reports whether signature is a valid HMAC of the
// payload under the webhook's secret. A webhook with no secret configured
// verifies nothing, so it is rejected rather than falling back to a shared one.
func ValidateWebhookSignature(payload []byte, secret string, signature string) bool {
	if secret == "" {
		return false
	}
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(signature), []byte(expected))
}
