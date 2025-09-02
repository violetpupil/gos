package os

import (
	"errors"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/gos/std/syscall"
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
