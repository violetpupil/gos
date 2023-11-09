package bytes

import (
	"bytes"
	"io"

	"github.com/sirupsen/logrus"
)

type (
	// 字节缓存
	// Bytes() 未读字节切片
	// WriteString() 将字符串写入缓存
	// If the buffer becomes too large, WriteString will panic with ErrTooLarge.
	Buffer = bytes.Buffer
)

var (
	// 创建字节缓存
	NewBuffer = bytes.NewBuffer
)

// NewBufferReader 从io.Reader构建bytes.Buffer
func NewBufferReader(src io.Reader) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, src)
	if err != nil {
		logrus.Errorln("io copy error", err)
		return nil, err
	}
	return buf, nil
}
