package errors

import (
	"errors"

	"github.com/violetpupil/components/lib/logrus"
)

// Unwrap 获取底层错误并打印类型
func Unwrap(err error) error {
	err = errors.Unwrap(err)
	if err != nil {
		logrus.Infof("unwrap %T %v", err, err)
	}
	return err
}
