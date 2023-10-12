package excelize

import (
	"fmt"
	"testing"
)

func TestNewFile(t *testing.T) {
	err := NewFile([][]any{
		{"姓名", "年龄"},
		{"张三", 18},
	}, "tmp")
	fmt.Println(err)
}
