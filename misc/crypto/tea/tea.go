// tea加密
// 密文为16进制字符串，每次加密得到的密文都不相同
// 明文越长，密文越长
// https://baike.baidu.com/item/TEA加密算法
// https://zh.wikipedia.org/zh-hans/微型加密算法
package tea

import (
	"bytes"
	"errors"
	"fmt"
)

const (
	// 密钥长度，只有前16个字节有效
	KeyLen = 16
	// 明文填充0字节长度
	// 设定为6，使密文长度最小为32
	ZeroLen = 6
	// 加密迭代次数
	Rounds = 16
)

// 迭代累加常量
var Delta uint32 = 0x9e3779b9

var (
	// 密钥无效
	ErrKey = errors.New("key invalid")
	// 加密块长度不足
	ErrBlockLen = errors.New("insufficient block length")
	// 密文无效
	ErrCipher = errors.New("cipher invalid")
)

// PadKey 填充密钥
func PadKey(key string) string {
	if len(key) >= KeyLen {
		return key
	}
	pad := bytes.Repeat([]byte{0}, KeyLen-len(key))
	key = fmt.Sprintf("%s_%s", key, string(pad))
	return key
}

// Swap 交换字节 0x01020304 -> 0x04030201
func Swap(x uint32) uint32 {
	x = x<<24&0xff000000 | x<<8&0x00ff0000 | x>>8&0x0000ff00 | x>>24&0x000000ff
	return x
}
