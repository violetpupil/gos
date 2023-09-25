package errors

import (
	"errors"
	"strings"

	"github.com/sirupsen/logrus"
)

// Unwrap 获取底层错误并打印类型
func Unwrap(err error) error {
	err = errors.Unwrap(err)
	if err != nil {
		logrus.Infof("unwrap %T %v", err, err)
	}
	return err
}

// Errors error切片
type Errors []error

// Error 将error用|拼接成字符串
func (es Errors) Error() string {
	var s string
	for _, e := range es {
		s += e.Error() + "|"
	}
	return strings.TrimSuffix(s, "|")
}
