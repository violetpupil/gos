package types

import "math"

// Range 生成 min~max 切片
func Range(min int, max int) []int {
	r := make([]int, 0)
	for i := min; i < max+1; i++ {
		r = append(r, i)
	}
	return r
}

// Hundred 获取百位数字
func Hundred(n int) int {
	return n / 100 % 10
}

// HundredThousand 获取十万位数字
func HundredThousand(n int) int {
	return n / int(math.Pow10(5)) % 10
}

// HundredMillion 获取亿位数字
func HundredMillion(n int) int {
	return n / int(math.Pow10(8)) % 10
}
