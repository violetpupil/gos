package reflect

import (
	"errors"
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

// FieldNames 获取结构体所有字段名
// v 是结构体或结构体指针
func FieldNames(v any) ([]string, error) {
	r := make([]string, 0)
	v = Pointer(v)

	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("not struct type")
	}
	for i := 0; i < t.NumField(); i++ {
		r = append(r, t.Field(i).Name)
	}
	return r, nil
}
