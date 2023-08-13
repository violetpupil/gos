package types

// Range 生成 min~max 切片
func Range(min int, max int) []int {
	r := make([]int, 0)
	for i := min; i < max+1; i++ {
		r = append(r, i)
	}
	return r
}
