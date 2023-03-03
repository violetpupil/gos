package crypto

import "encoding/hex"

// DecryptStr 解密字符串，密钥长度不固定，只有前16个字节有效
func DecryptStr(cipher, key string) (string, error) {
	key = PadKey(key)
	cipherB, err := hex.DecodeString(cipher)
	if err != nil {
		return "", err
	}
	plain := DecryptByte(cipherB, []byte(key))
	return string(plain), nil
}

// DecryptByte 解密字节
func DecryptByte(cipher, key []byte) []byte {
	// TODO
	return cipher
}

// Decrypt 解密
func Decrypt(v0, v1, k0, k1, k2, k3 uint32) (uint32, uint32) {
	v0, v1 = Swap(v0), Swap(v1)

	// 只使用后32位
	sum := Delta * Rounds
	for i := 0; i < Rounds; i++ {
		// + 优先级大于 ^
		v1 -= (v0 << 4) + k2 ^ v0 + sum ^ (v0 >> 5) + k3
		v0 -= (v1 << 4) + k0 ^ v1 + sum ^ (v1 >> 5) + k1
		sum -= Delta
	}
	v0, v1 = Swap(v0), Swap(v1)
	return v0, v1
}
