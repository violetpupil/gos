// 切片
// 逐个添加元素，扩容后的容量首次为2^3，之后指数递增
package types

import (
	"sort"
)

// Reverse 反转切片
func Reverse(x any) {
	sort.Slice(x, func(i, j int) bool { return true })
}

// Safe 切片索引是否安全，不会导致panic
func Safe(start, end, cap int) bool {
	return 0 <= start && start <= end && end <= cap
}

// Div 将切片按顺序转为二维切片
func Div[T any](src []T, size int) [][]T {
	dst := make([][]T, 0)
	for i := 0; i < len(src); i += size {
		if i+size >= len(src) {
			dst = append(dst, src[i:])
		} else {
			dst = append(dst, src[i:i+size])
		}
	}
	return dst
}
