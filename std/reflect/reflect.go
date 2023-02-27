package reflect

import "reflect"

// Value 接口值
// 零值方法调用会异常，除了
// IsValid 返回 false
// Kind 返回 Invalid
// String 返回 "<invalid Value>"
type Value = reflect.Value

// Elem 提取指针指向的值
func Elem(i any) any {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		// 指针是nil，Elem会返回零值
		return v.Elem().Interface()
	}
	return i
}
