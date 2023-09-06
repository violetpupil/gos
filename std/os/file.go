package os

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/gos/std/fs"
)

// 函数
var (
	// ReadFile 读取文件字节
	ReadFile = os.ReadFile
	// Create 创建文件，如果文件存在，先清空
	Create = os.Create
	// IsNotExist 是否错误为不存在
	IsNotExist = os.IsNotExist
)

type (
	// File represents an open file descriptor.
	// Close() 关闭文件
	File = os.File
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

// Exist returns true if a file or directory exists.
func Exist(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	} else {
		logrus.Error("stat error ", err)
		return false, err
	}
}

// HasContent 文件是否存在，并且有内容
func HasContent(name string) (bool, error) {
	info, err := os.Stat(name)
	if err == nil {
		return info.Size() > 0, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	} else {
		logrus.Error("stat error ", err)
		return false, err
	}
}

// WriteFile 写文件
// 文件不存在，创建文件，权限设置为可读可写
// 文件存在，先清空文件后写入，权限不变
// 不会自动创建文件夹
func WriteFile(name string, data []byte) error {
	return os.WriteFile(name, data, 0666)
}

// Rename 将文件夹所有文件重命名
func Rename(dir string, fn func(string) string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		logrus.Errorln("read dir error", err)
		return err
	}
	for _, f := range files {
		old := filepath.Join(dir, f.Name())
		new := filepath.Join(dir, fn(f.Name()))
		err := os.Rename(old, new)
		if err != nil {
			logrus.Errorln("rename error", err)
			return err
		}
	}
	return nil
}
