package reflect

import (
	"fmt"
	"reflect"
)

// 具体接口类型
var (
	TypeStringer = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	TypeError    = reflect.TypeOf((*error)(nil)).Elem()
)

// Value 接口值
// 零值方法调用会异常，除了
// IsValid 返回 false
// Kind 返回 Invalid
// String 返回 "<invalid Value>"
type Value = reflect.Value

// Pointer 提取指针指向的值
func Pointer(i any) any {
	if i == nil {
		return nil
	}

	v := reflect.ValueOf(i)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		// 指针是nil，Elem会返回零值
		v = v.Elem()
	}
	return v.Interface()
}
