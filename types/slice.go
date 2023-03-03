// 切片
// 逐个添加元素，扩容后的容量首次为2^3，之后指数递增
package types

import "sort"

// Reverse 反转切片
func Reverse(x any) {
	sort.Slice(x, func(i, j int) bool { return true })
}
