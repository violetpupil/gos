package rand

import "crypto/rand"

// RandBytes 生成指定长度的随机字节
func RandBytes(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
