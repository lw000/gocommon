package tydes

import (
	"crypto/des"
	"encoding/hex"
	"errors"
	"github.com/lw000/gocommon/auth"
)

type DesUtil struct {
	key []byte
}

func New(key []byte) *DesUtil {
	return &DesUtil{key: key}
}

func (d DesUtil) ecbEncrypt(data []byte, fillMode int) ([]byte, error) {
	if len(d.key) == 0 {
		return nil, errors.New("key is empty")
	}

	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}

	var key []byte

	if len(d.key) > 8 {
		key = d.key[:8]
	} else {
		key = d.key
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bs := block.BlockSize()
	switch fillMode {
	case 0:
		data = tyauth.ZeroPadding(data, bs)
	case 5:
		data = tyauth.Pkcs5Padding(data, bs)
	case 7:
		data = tyauth.Pkcs5Padding(data, bs)
	default:
		return nil, errors.New("nonsupport")
	}

	if len(data)%bs != 0 {
		return nil, errors.New("input not full blocks")
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}

	return out, nil
}

func (d DesUtil) ecbDecrypt(data []byte, fillMode int) ([]byte, error) {
	if len(d.key) == 0 {
		return nil, errors.New("key is empty")
	}

	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}

	var (
		newKey []byte
	)

	if len(d.key) > 8 {
		newKey = d.key[:8]
	} else {
		newKey = d.key
	}

	block, err := des.NewCipher(newKey)
	if err != nil {
		return nil, err
	}

	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}

	switch fillMode {
	case 0:
		out = tyauth.ZeroUnPadding(out)
	case 5:
		out = tyauth.Pkcs5UnPadding(out)
	case 7:
		out = tyauth.Pkcs5UnPadding(out)
	default:
		return nil, errors.New("nonsupport")
	}

	return out, nil
}

func EcbEncrypt(key []byte, data []byte) (string, error) {
	d := New(key)
	data, err := d.ecbEncrypt(data, 0)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(data), nil
}

func EcbDecrypt(key []byte, str string) ([]byte, error) {
	d := New(key)
	data, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return d.ecbDecrypt(data, 0)
}
