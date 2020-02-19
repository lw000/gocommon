package tyrsa

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/lw000/gocommon/utils"
)

/*
RSA算法本身要求加密内容也就是明文长度m必须0<m<密钥长度n。
如果小于这个长度就需要进行padding，因为如果没有padding，就无法确定解密后内容的真实长度，字符串之类的内容问题还不大，
以0作为结束符，但对二进制数据就很难，因为不确定后面的0是内容还是内容结束符。而只要用到padding，那么就要占用实际的明文长度，
于是实际明文长度需要减去padding字节长度。我们一般使用的padding标准有NoPPadding、OAEPPadding、PKCS1Padding等，
其中PKCS#1建议的padding就占用了11个字节。

这样，对于1024长度的密钥。128字节（1024bits）-减去11字节正好是117字节，但对于RSA加密来讲，padding也是参与加密的，
所以，依然按照1024bits去理解，但实际的明文只有117字节了。
*/

type RSAUtil struct {
	publicKey  []byte
	privateKey []byte
}

func NewRSA(publicKey []byte, privateKey []byte) *RSAUtil {
	return &RSAUtil{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (r *RSAUtil) pubKey() (*rsa.PublicKey, error) {
	block, _ := pem.Decode(r.publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	// pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// return pub, nil

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)

	return pub, nil
}

func (r *RSAUtil) priKey() (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(r.privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	priInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pri := priInterface.(*rsa.PrivateKey)

	return pri, nil
}

func (r *RSAUtil) EncryptByPublicKey(originData []byte) ([]byte, error) {
	if len(originData) == 0 {
		return nil, errors.New("encrypt originData is empty")
	}

	pub, err := r.pubKey()
	if err != nil {
		return nil, err
	}

	partLen := pub.N.BitLen()/8 - 11
	chunks := tyutils.Split(originData, partLen)
	buffer := bytes.NewBuffer(nil)
	for _, chunk := range chunks {
		var buf []byte
		buf, err = rsa.EncryptPKCS1v15(rand.Reader, pub, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(buf)
	}
	return buffer.Bytes(), nil
}

func (r *RSAUtil) DecryptByPrivateKey(originData []byte) ([]byte, error) {
	if len(originData) == 0 {
		return nil, errors.New("decrypt originData is empty")
	}

	priv, err := r.priKey()
	if err != nil {
		return nil, err
	}

	partLen := priv.N.BitLen() / 8
	chunks := tyutils.Split(originData, partLen)
	buffer := bytes.NewBuffer(nil)
	for _, chunk := range chunks {
		var buf []byte
		buf, err = rsa.DecryptPKCS1v15(rand.Reader, priv, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(buf)
	}
	return buffer.Bytes(), nil
}

func (r *RSAUtil) Sign(originData []byte, hash crypto.Hash) (string, error) {
	if len(originData) == 0 {
		return "", errors.New("sign originData is empty")
	}

	priv, err := r.priKey()
	if err != nil {
		return "", err
	}

	h := hash.New()
	h.Write(originData)
	hashed := h.Sum(nil)
	originData, err = rsa.SignPKCS1v15(rand.Reader, priv, hash, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(originData), nil
}

func (r *RSAUtil) Verify(originData []byte, signData string, hash crypto.Hash) error {
	if len(originData) == 0 {
		return errors.New("verify originData is empty")
	}

	if signData == "" {
		return errors.New("verify signData is empty")
	}

	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}

	pub, err := r.pubKey()
	if err != nil {
		return err
	}

	h := hash.New()
	h.Write([]byte(originData))
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(pub, hash, hashed, sign)
}
