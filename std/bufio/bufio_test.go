package bufio

import (
	"fmt"
	"testing"
)

// 测试无效，会立即结束
func TestScan(t *testing.T) {
	Scan(func(s string, err error) bool {
		if err == nil {
			fmt.Println(s)
		} else {
			fmt.Println("scan error ", err)
		}
		return false
	})
}
