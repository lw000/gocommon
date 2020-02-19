package tyauth

import (
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"hash"
	"io"
)

func Hex(src []byte) string {
	return hex.EncodeToString(src)
}

func Hmac(s string, key string) (string, error) {
	h := hmac.New(func() hash.Hash {
		return crypto.SHA256.New()
	}, []byte(key))
	n, err := io.WriteString(h, s)
	if err != nil {
		return "", err
	}
	if n > 0 {

	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
