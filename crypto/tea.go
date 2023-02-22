// tea加密
// 密文为16进制32位字符串，每次加密得到的密文都不相同
package crypto

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

// Encrypt 加密字符串
func Encrypt(plain, key string) string {
	key = PadKey(key)
	cipherB := encrypt([]byte(plain), []byte(key))
	cipher := hex.EncodeToString(cipherB)
	return cipher
}

// Decrypt 解密字符串
func Decrypt(cipher, key string) (string, error) {
	key = PadKey(key)
	cipherB, err := hex.DecodeString(cipher)
	if err != nil {
		return "", err
	}
	plain := decrypt(cipherB, []byte(key))
	return string(plain), nil
}

// PadKey 填充密钥，结果最少16位
func PadKey(key string) string {
	if len(key) >= 16 {
		return key
	}
	pad := bytes.Repeat([]byte{0}, 16-len(key))
	key = fmt.Sprintf("%s_%s", key, string(pad))
	return key
}

// encrypt 加密字节
func encrypt(plain, key []byte) []byte {
	// TODO
	return plain
}

// decrypt 解密字节
func decrypt(cipher, key []byte) []byte {
	// TODO
	return cipher
}
