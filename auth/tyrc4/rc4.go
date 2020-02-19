package tyrc4

import (
	"crypto/rc4"
	"encoding/base64"
	"errors"
)

func RC4(key []byte, data []byte) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key is empty")
	}

	if len(data) == 0 {
		return "", errors.New("data is empty")
	}

	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return "", err
	}

	dst := make([]byte, len(data))
	cipher.XORKeyStream(dst, data)

	return base64.StdEncoding.EncodeToString(dst), nil
}
