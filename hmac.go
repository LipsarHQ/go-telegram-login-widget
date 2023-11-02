package telegramloginwidget

import (
	"crypto/hmac"
	"hash"
)

func hashHMAC(data, key []byte, fn func() hash.Hash) []byte {
	h := hmac.New(fn, key)
	_, _ = h.Write(data) //nolint:errcheck // No need to check error.

	return h.Sum(nil)
}
