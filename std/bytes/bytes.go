package bytes

import (
	"bytes"
	"io"

	"github.com/sirupsen/logrus"
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
