package bufio

import "bufio"

type (
	// Reader 将io.Reader的内容读到内存
	Reader = bufio.Reader
)

var (
	// NewReader 创建Reader，默认buffer是4096字节
	NewReader = bufio.NewReader
)
