package rand

import (
	"crypto/rand"

	"github.com/violetpupil/components/lib/logrus"
)

// RandBytes 生成指定长度的随机字节
func RandBytes(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		logrus.Error("rand read error ", err)
		return nil, err
	}
	return b, nil
}

// RandByte 生成随机字节
func RandByte() (byte, error) {
	bs, err := RandBytes(1)
	if err != nil {
		logrus.Error("rand bytes error ", err)
		return 0, err
	}
	return bs[0], nil
}
