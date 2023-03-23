package os

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// 函数
var (
	// ReadFile 读取文件字节
	ReadFile = os.ReadFile
)

// 结构体
type (
	// SyscallError 系统调用错误
	SyscallError = os.SyscallError
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
}

type FileInfo struct {
	Name    string      // 文件名
	Size    int64       // 文件字节数
	Mode    os.FileMode // 文件模式
	ModTime time.Time   // 更新时间
	IsDir   bool        // 是否为目录
}

// Stat 获取文件信息
func Stat(name string) (*FileInfo, error) {
	info, err := os.Stat(name)
	if err != nil {
		logrus.Error("stat error ", err)
		return nil, err
	}
	fi := &FileInfo{
		Name:    info.Name(),
		Size:    info.Size(),
		Mode:    info.Mode(),
		ModTime: info.ModTime(),
		IsDir:   info.IsDir(),
	}
	return fi, nil
}

// WriteFile 写文件
// 文件不存在，创建文件，权限设置为可读可写
// 文件存在，先清空文件后写入，权限不变
func WriteFile(name string, data []byte) error {
	return os.WriteFile(name, data, 0666)
}
