// https://gist.github.com/hothero/7d085573f5cb7cdb5801d7adcf66dcf3
package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/sirupsen/logrus"
)

// Encrypt aes加密
func Encrypt(key, iv, plain []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		logrus.Errorln("new cipher error", err)
		return nil, err
	}

	var c []byte
	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(c, plain)
	return c, nil
}

// EncryptS aes加密字符串，返回base64 url编码的密文
func EncryptS(key, iv []byte, plain string) (string, error) {
	cipher, err := Encrypt(key, iv, []byte(plain))
	if err != nil {
		logrus.Errorln("encrypt error", err)
		return "", err
	}
	c := base64.URLEncoding.EncodeToString(cipher)
	return c, nil
}
