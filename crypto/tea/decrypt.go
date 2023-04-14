package tea

import (
	"encoding/hex"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/spec/types"
	"github.com/violetpupil/components/std/binary"
)

// DecryptStr 解密字符串，密钥长度不固定，只有前16个字节有效
func DecryptStr(cipher, key string) (string, error) {
	key = PadKey(key)
	cipherB, err := hex.DecodeString(cipher)
	if err != nil {
		logrus.Error("decode string error ", err)
		return "", err
	}
	plain, err := DecryptByte(cipherB, []byte(key))
	return string(plain), err
}

// DecryptByte 解密字节
func DecryptByte(cipher, key []byte) ([]byte, error) {
	ks := binary.SplitUint32(binary.BigEndian, key)
	if len(ks) < 4 {
		return nil, ErrKey
	}
	k0, k1, k2, k3 := ks[0], ks[1], ks[2], ks[3]

	// 分块明文存储当前块明文和上一块密文异或结果
	// 分块密文存储当前块密文
	var plainB, cipherB [2]uint32
	plain := make([]byte, 0)
	for i := 0; i < len(cipher); i += 8 {
		// 8个字节为一块
		block := cipher[i : i+8]
		vs := binary.SplitUint32(binary.LittleEndian, block)
		if len(vs) < 2 {
			return nil, ErrBlockLen
		}

		// 一块分为两个uint32组
		// 相同的异或操作还原明文
		v0, v1 := vs[0], vs[1]
		cipherTmp := [2]uint32{v0, v1}
		v0 ^= plainB[0]
		v1 ^= plainB[1]

		v0, v1 = Decrypt(v0, v1, k0, k1, k2, k3)
		plainB = [2]uint32{v0, v1}
		v0 ^= cipherB[0]
		v1 ^= cipherB[1]
		cipherB = cipherTmp

		block = binary.ComUint32(binary.LittleEndian, v0, v1)
		plain = append(plain, block...)
	}
	// 去掉明文填充字节
	plain, err := UnPadPlain(plain)
	return plain, err
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

// UnPadPlain 还原明文 header+pad+salt+plain+zeros
func UnPadPlain(plain []byte) ([]byte, error) {
	// 密文无效，导致明文长度不足
	if len(plain) < 1 {
		return nil, ErrCipher
	}

	// header后3位代表pad长度
	padLen := plain[0] & 7
	// header为1个字节，salt为2个字节
	start := 1 + 2 + padLen
	end := len(plain) - ZeroLen
	if !types.Safe(int(start), end, cap(plain)) {
		return nil, ErrCipher
	}

	plain = plain[start:end]
	return plain, nil
}
