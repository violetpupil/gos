package reflect

import (
	"errors"
	"reflect"

	"github.com/sirupsen/logrus"
)

// Pointer 提取指针指向的值
// 可以传入非指针，原样返回
// 发现nil指针，返回nil
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
// v 是结构体或结构体指针，指针不能为nil
// 因为嵌入结构体的字段也会解析，嵌入结构体指针也不能为nil
func FieldNames(v any) ([]string, error) {
	r := make([]string, 0)
	// 去掉可能的指针
	v = Pointer(v)
	if v == nil {
		return nil, errors.New("cannot be nil")
	}

	// 确保为结构体
	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("not struct type")
	}

	// 解析字段名
	for i := 0; i < typ.NumField(); i++ {
		ft := typ.Field(i)
		fv := val.Field(i)
		if !ft.Anonymous {
			r = append(r, ft.Name)
		} else {
			// 递归处理嵌入结构体
			var err error
			r, err = FieldNames(fv.Interface())
			if err != nil {
				logrus.Errorln("field names error", err)
				return nil, err
			}
		}
	}
	return r, nil
}
