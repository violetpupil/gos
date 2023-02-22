package crypto

import (
	"encoding/binary"
	"strings"
)

// Format 数据格式
type Format struct {
	ByteOrder binary.ByteOrder
	Format    string
}

// NewFormat 将格式字符串转为对象
// 可以使用第一个字符指定字节序，不指定则默认大端
func NewFormat(src string) *Format {
	f := new(Format)
	if strings.HasPrefix(src, "<") {
		f.ByteOrder = binary.LittleEndian
		f.Format = src[1:]
	} else if strings.HasPrefix(src, ">") {
		f.ByteOrder = binary.BigEndian
		f.Format = src[1:]
	} else {
		f.ByteOrder = binary.BigEndian
		f.Format = src
	}
	return f
}

// Pack 将数据打包成字节串
func (f *Format) Pack(data ...any) ([]byte, error) {
	// TODO
	return nil, nil
}

// Unpack 将字节串解包
func (f *Format) Unpack() {
	// TODO
}
