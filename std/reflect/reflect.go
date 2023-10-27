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
	// 直到非指针或nil指针
	// nil指针，Elem会返回零值 <invalid reflect.Value>
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	// 处理nil指针
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return nil
	}
	return v.Interface()
}

// FieldNames 获取结构体所有字段名
// v 是结构体或结构体指针
func FieldNames(v any) ([]string, error) {
	r := make([]string, 0)
	v = Pointer(v)
	if v == nil {
		return nil, errors.New("cannot be nil")
	}

	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("not struct type")
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		r = append(r, f.Name)
	}
	return r, nil
}
