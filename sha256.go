package telegramloginwidget

import (
	"crypto/sha256"
)

func hashSHA256(data []byte) []byte {
	h := sha256.New()
	_, _ = h.Write(data) //nolint:errcheck // No need to check error.

	return h.Sum(nil)
}
