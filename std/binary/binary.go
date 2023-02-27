package binary

import "encoding/binary"

var (
	// 小端序 低位字节在低地址
	LittleEndian = binary.LittleEndian
	// 大端序 高位字节在低地址
	BigEndian = binary.BigEndian
)
