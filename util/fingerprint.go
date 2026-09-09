package util

import (
	"crypto/md5"
	"fmt"
)

// FingerprintCredentials returns a stable fingerprint for a set of SMTP
// credentials, so duplicate sending profiles can be detected without keeping
// the credentials themselves around.
func FingerprintCredentials(username string, password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(username+":"+password)))
}
