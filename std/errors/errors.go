// An error e wraps another error if e's type has one of the methods
// Unwrap() error
// Unwrap() []error
package errors

import (
	"errors"

	"github.com/sirupsen/logrus"
)

var (
	// 获取 error 树中的指定 error
	As = errors.As
	// error 树是否有指定 error
	Is = errors.Is
	// 包裹多个 error 成一个
	// 返回的 error 实现 Unwrap() []error 方法
	// 返回的 error 字符串为每个 error 用换行连接
	// 传入的 error 都是 nil，返回 nil
	Join = errors.Join
	// 使用字符串创建 error
	New = errors.New
)

// Unwrap 获取底层错误并打印类型
func Unwrap(err error) error {
	// 调用 err 的 Unwrap() error 方法
	// err 没有实现对应方法，返回 nil
	err = errors.Unwrap(err)
	if err != nil {
		logrus.Infof("unwrap %T %v", err, err)
	}
	return err
}
