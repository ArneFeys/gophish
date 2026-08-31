package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// FingerprintCredentials returns a stable fingerprint for a set of SMTP
// credentials, so duplicate sending profiles can be detected without keeping
// the credentials themselves around. The fingerprint is keyed, so it cannot be
// brute-forced back to the password the way a bare digest can.
func FingerprintCredentials(key []byte, username string, password string) string {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(username + ":" + password))
	return hex.EncodeToString(mac.Sum(nil))
}
