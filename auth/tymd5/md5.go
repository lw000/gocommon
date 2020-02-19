package tymd5

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
)

func MD5(str []byte) (string, error) {
	if len(str) == 0 {
		return "", errors.New("str is empty")
	}

	h := md5.New()
	_, err := h.Write(str)
	if err != nil {
		return "", err
	}

	s := hex.EncodeToString(h.Sum(nil))
	return s, nil
}

func MD5String(str string) (string, error) {
	return MD5([]byte(str))
}

func MD5File(name string) (string, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}
	return MD5(data)
}
