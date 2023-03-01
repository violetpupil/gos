package binary

import "encoding/binary"

var (
	// 小端序 低位字节在低地址
	LittleEndian = binary.LittleEndian
	// 大端序 高位字节在低地址
	BigEndian = binary.BigEndian
)

// ByteOrder 字节序
type ByteOrder interface {
	binary.ByteOrder
	binary.AppendByteOrder
}

// ComUint32 将uint32转为字节串
func ComUint32(bo ByteOrder, data ...uint32) []byte {
	bs := make([]byte, 0)
	for _, v := range data {
		bs = bo.AppendUint32(bs, v)
	}
	return bs
}

// SplitUint32 将字节串转为uint32
// 如果字节串长度不是4的倍数，最后的字节会被丢掉
func SplitUint32(bo ByteOrder, bs []byte) []uint32 {
	data := make([]uint32, 0)
	for len(bs) >= 4 {
		// 将前4个字节转为uint32
		data = append(data, bo.Uint32(bs))
		bs = bs[4:]
	}
	return data
}
