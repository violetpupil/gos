package bufio

import (
	"bufio"
	"os"
)

type (
	// Reader 将io.Reader的内容读到内存
	Reader = bufio.Reader
	// 读取行分隔文本
	// 使用Scan扫描，使用Text获取扫描结果，使用Err获取错误
	Scanner = bufio.Scanner
)

var (
	// NewReader 创建Reader，默认buffer是4096字节
	NewReader = bufio.NewReader
)

// Scan 循环读取终端输入
// 参数是处理函数，调用有两种情况
// 1 单行文本，nil
// 2 ""，导致循环结束的错误
// ctrl-c 会使标准输入结束，此时不会回调处理函数
// 处理函数返回是否要结束循环
func Scan(f func(string, error) bool) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if f(s.Text(), nil) {
			return
		}
	}

	// 输入结束
	if s.Err() == nil {
		return
	}
	// 发生错误
	f("", s.Err())
}
