// https://gorm.io/docs/security.html
package gorm

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/gos/std/reflect"
	"golang.org/x/exp/slices"
)

// OrderCheck 排序字段检查
// v 是结构体或结构体指针，指针不能为nil，嵌入结构体指针也不能为nil
func OrderCheck(v, any, s string) (bool, error) {
	fields, err := reflect.FieldNames(v)
	if err != nil {
		logrus.Errorln("field names error", err)
		return false, err
	}

	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, ",", " ")
	parts := strings.Split(s, " ")
	for _, p := range parts {
		if p == "" || p == "asc" || p == "desc" {
			continue
		}
		if slices.Contains(fields, p) {
			continue
		}
		return false, nil
	}
	return true, nil
}
