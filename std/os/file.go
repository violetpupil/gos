package os

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/gos/std/fs"
)

var (
	// 文件操作权限
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY = os.O_RDONLY // 只读
	O_WRONLY = os.O_WRONLY // 只写
	O_RDWR   = os.O_RDWR   // 读写
	// The remaining values may be or'ed in to control behavior.
	O_APPEND = os.O_APPEND // append data to the file when writing.
	O_CREATE = os.O_CREATE // create a new file if none exists.
	O_EXCL   = os.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   = os.O_SYNC   // open for synchronous I/O.
	O_TRUNC  = os.O_TRUNC  // truncate regular writable file when opened.
)

// 函数
var (
	// ReadFile 读取文件字节
	ReadFile = os.ReadFile
	// Open 打开文件，仅有读权限
	Open = os.Open
	// Create 创建文件，如果文件存在，先清空
	Create = os.Create
	// Remove 删除文件或空文件夹，路径不存在则报错
	Remove = os.Remove
	// RemoveAll 删除文件或整个文件夹，路径不存在则忽略
	RemoveAll = os.RemoveAll
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

// Size 获取文件大小字节数
func Size(name string) (int64, error) {
	info, err := os.Stat(name)
	if err != nil {
		logrus.Error("stat error ", err)
		return 0, err
	}
	return info.Size(), nil
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

// Mkdir 创建单级文件夹，存在返回nil
func Mkdir(name string) error {
	err := os.Mkdir(name, 0666)
	if err == nil || os.IsExist(err) {
		return nil
	} else {
		logrus.Errorln("mkdir error", err)
		return err
	}
}

// Mkdir 创建多级文件夹，存在返回nil
func MkdirAll(path string) error {
	err := os.MkdirAll(path, 0666)
	return err
}

// Clear 清空文件
func Clear(name string) error {
	return os.Truncate(name, 0)
}
