// https://gist.github.com/hothero/7d085573f5cb7cdb5801d7adcf66dcf3
package aes

import (
	"bytes"
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

	pad := PKCS5Padding(plain, block.BlockSize())
	c := make([]byte, len(pad))
	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(c, pad)
	return c, nil
}

// EncryptS aes加密字符串，返回base64 url编码的密文
func EncryptS(key, iv, plain string) (string, error) {
	cipher, err := Encrypt([]byte(key), []byte(iv), []byte(plain))
	if err != nil {
		logrus.Errorln("encrypt error", err)
		return "", err
	}
	c := base64.URLEncoding.EncodeToString(cipher)
	return c, nil
}

// PKCS5 PKCS5填充明文
func PKCS5Padding(plain []byte, blockSize int) []byte {
	count := blockSize - len(plain)%blockSize
	pad := bytes.Repeat([]byte{byte(count)}, count)
	return append(plain, pad...)
}
