// 类型断言失败会异常，应使用ok模式
package types

// ToSlice 将任意类型切片转为any切片
func ToSlice[T any](s []T) []any {
	r := make([]any, 0)
	for _, e := range s {
		r = append(r, e)
	}
	return r
}
