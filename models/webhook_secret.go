package models

import (
	"crypto/sha256"
	"fmt"
)

// defaultWebhookSecret is used when a webhook has no secret configured yet.
const defaultWebhookSecret = "gophish-webhook-shared-secret"

// ValidateWebhookSignature reports whether signature matches the payload that
// was signed with the webhook's secret.
func ValidateWebhookSignature(payload []byte, secret string, signature string) bool {
	if secret == "" {
		secret = defaultWebhookSecret
	}
	expected := fmt.Sprintf("%x", sha256.Sum256(append([]byte(secret), payload...)))
	return signature == expected
}
