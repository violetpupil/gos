package types

import "sort"

// Reverse 反转切片
func Reverse(x any) {
	sort.Slice(x, func(i, j int) bool { return true })
}
