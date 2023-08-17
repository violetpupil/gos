package rand

import "math/rand"

// RandomInt 返回[start,end)随机数，end<=start会panic
func RandomInt(start, end int) int {
	return rand.Intn(end-start) + start
}

// RandomSlice 随机切片元素
func RandomSlice[T any](slice []T) T {
	return slice[RandomInt(0, len(slice))]
}
