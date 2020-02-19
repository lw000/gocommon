package tyaes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"github.com/lw000/gocommon/auth"
)

type AesUtil struct {
	key []byte
}

func New(key []byte) *AesUtil {
	return &AesUtil{key: key}
}

func (a AesUtil) cbcEncrypt(data []byte, fillMode int) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	switch fillMode {
	case 0:
		data = tyauth.ZeroPadding(data, blockSize)
	case 5:
		data = tyauth.Pkcs5Padding(data, blockSize)
	case 7:
		data = tyauth.Pkcs5Padding(data, blockSize)
	default:
		return nil, errors.New("nonsupport")
	}

	blockMode := cipher.NewCBCEncrypter(block, a.key[:blockSize])
	originData := make([]byte, len(data))
	blockMode.CryptBlocks(originData, data)

	return originData, nil
}

func (a AesUtil) cbcDecrypt(data []byte, fillMode int) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, a.key[:blockSize])

	var originData []byte
	originData = make([]byte, len(data))
	blockMode.CryptBlocks(originData, []byte(data))
	switch fillMode {
	case 0:
		originData = tyauth.ZeroUnPadding(originData)
	case 5:
		originData = tyauth.Pkcs5UnPadding(originData)
	case 7:
		originData = tyauth.Pkcs5UnPadding(originData)
	default:
		return nil, errors.New("nonsupport")
	}

	return originData, nil
}

func Encrypt(key []byte, data []byte) (string, error) {
	as := New(key)
	originData, err := as.cbcEncrypt(data, 5)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(originData), err
}

func Decrypt(key []byte, str string) ([]byte, error) {
	as := New(key)
	data, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return as.cbcDecrypt(data, 5)
}
