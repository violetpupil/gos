package os

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/std/fs"
)

// 函数
var (
	// ReadFile 读取文件字节
	ReadFile = os.ReadFile
	// Create 创建文件，如果文件存在，先清空
	Create = os.Create
)

// Stat 获取文件信息
func Stat(name string) (*fs.FileInfoS, error) {
	info, err := os.Stat(name)
	if err != nil {
		logrus.Error("stat error ", err)
		return nil, err
	}
	fi := &fs.FileInfoS{
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
