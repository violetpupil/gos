package rand

import "math/rand"

// RandomInt 返回[start,end)随机数，end<=start会panic
func RandomInt(start, end int) int {
	return rand.Intn(end-start) + start
}
