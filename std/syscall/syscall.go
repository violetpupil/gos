package syscall

import (
	"errors"
	"syscall"

	"github.com/sirupsen/logrus"
)

type (
	// 系统调用错误，因操作系统而异，0代表没有错误
	Errno = syscall.Errno
)

// LogErrno 打印syscall.Errno相关信息
// 如果不是syscall.Errno，直接返回
func LogErrno(err error) {
	var no syscall.Errno
	if !errors.As(err, &no) {
		return
	}

	logrus.WithFields(logrus.Fields{
		"Error":     no,
		"Timeout":   no.Timeout(),
		"Temporary": no.Temporary(),
	}).Errorf("%d", no)
}
