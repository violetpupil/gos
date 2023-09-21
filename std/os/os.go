package os

import (
	"errors"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/gos/std/syscall"
)

// 结构体
type (
	// SyscallError 系统调用错误
	SyscallError = os.SyscallError
)

var (
	// 将${var}或$var替换为环境变量或者空字符串
	ExpandEnv = os.ExpandEnv
)

var (
	// Args hold the command-line arguments, starting with the program name.
	Args = os.Args
)

// LogSyscallError 打印os.SyscallError相关信息
// 如果不是os.SyscallError，直接返回
func LogSyscallError(err error) {
	scErr := new(os.SyscallError)
	if !errors.As(err, &scErr) {
		return
	}

	errWrap := errors.Unwrap(scErr)
	logrus.WithFields(logrus.Fields{
		"Error":         scErr,
		"Timeout":       scErr.Timeout(),
		"ErrorWrapType": fmt.Sprintf("%T", errWrap),
	}).Errorf("%+v", *scErr)

	syscall.LogErrno(errWrap)
}

// Environ 打印所有环境变量 key=value
func Environ() {
	fmt.Println("all env var")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	fmt.Println("")
}

// Setenv 批量设置环境变量
func Setenv(kv map[string]string) error {
	for k, v := range kv {
		err := os.Setenv(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
