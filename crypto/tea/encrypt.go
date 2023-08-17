package tea

import (
	"bytes"
	"encoding/hex"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/gos/std/binary"
	"github.com/violetpupil/gos/std/crypto/rand"
)

// EncryptStr 加密字符串，密钥长度不固定，只有前16个字节有效
func EncryptStr(plain, key string) (string, error) {
	key = PadKey(key)
	cipherB, err := EncryptByte([]byte(plain), []byte(key))
	if err != nil {
		logrus.Error("encrypt byte error ", err)
		return "", err
	}
	cipher := hex.EncodeToString(cipherB)
	return cipher, nil
}

// EncryptByte 加密字节
func EncryptByte(plain, key []byte) ([]byte, error) {
	// 填充明文，最终总长度为8的倍数
	plain, err := PadPlain(plain)
	if err != nil {
		logrus.Error("pad plain error ", err)
		return nil, err
	}
	ks := binary.SplitUint32(binary.BigEndian, key)
	if len(ks) < 4 {
		return nil, ErrKey
	}
	k0, k1, k2, k3 := ks[0], ks[1], ks[2], ks[3]

	// 分块明文存储当前块明文和上一块密文异或结果
	// 分块密文存储当前块密文
	var plainB, cipherB [2]uint32
	cipher := make([]byte, 0)
	for i := 0; i < len(plain); i += 8 {
		// 8个字节为一块
		block := plain[i : i+8]
		vs := binary.SplitUint32(binary.LittleEndian, block)
		if len(vs) < 2 {
			return nil, ErrBlockLen
		}

		// 一块分为两个uint32组
		v0, v1 := vs[0], vs[1]
		v0 ^= cipherB[0]
		v1 ^= cipherB[1]
		plainTmp := [2]uint32{v0, v1}

		v0, v1 = Encrypt(v0, v1, k0, k1, k2, k3)
		v0 ^= plainB[0]
		v1 ^= plainB[1]
		plainB = plainTmp
		cipherB = [2]uint32{v0, v1}

		block = binary.ComUint32(binary.LittleEndian, v0, v1)
		cipher = append(cipher, block...)
	}
	return cipher, nil
}

// Encrypt 加密
func Encrypt(v0, v1, k0, k1, k2, k3 uint32) (uint32, uint32) {
	v0, v1 = Swap(v0), Swap(v1)

	var sum uint32
	for i := 0; i < Rounds; i++ {
		// 只使用后32位
		sum += Delta
		// + 优先级大于 ^
		v0 += (v1 << 4) + k0 ^ v1 + sum ^ (v1 >> 5) + k1
		v1 += (v0 << 4) + k2 ^ v0 + sum ^ (v0 >> 5) + k3
	}
	v0, v1 = Swap(v0), Swap(v1)
	return v0, v1
}

// PadPlain 填充明文 header+pad+salt+plain+zeros
func PadPlain(plain []byte) ([]byte, error) {
	// 计算填充长度，使最终总长度为8的倍数
	// header为1个字节，salt为2个字节
	total := 1 + 2 + len(plain) + ZeroLen
	padLen := total % 8
	if padLen > 0 {
		padLen = 8 - padLen
	}

	// header前5位随机，后3位代表填充长度
	header, err := rand.RandByte()
	if err != nil {
		logrus.Error("gen header error ", err)
		return nil, err
	}
	header = header&0xf8 | byte(padLen)
	// 填充随机字节
	pad, err := rand.RandBytes(padLen)
	if err != nil {
		logrus.Error("gen pad error ", err)
		return nil, err
	}
	salt, err := rand.RandBytes(2)
	if err != nil {
		logrus.Error("gen salt error ", err)
		return nil, err
	}
	// 填充0字节，加长密文长度
	zeros := bytes.Repeat([]byte{0}, ZeroLen)

	bs := make([]byte, 0)
	bs = append(bs, header)
	bs = append(bs, pad...)
	bs = append(bs, salt...)
	bs = append(bs, plain...)
	bs = append(bs, zeros...)
	return bs, nil
}
