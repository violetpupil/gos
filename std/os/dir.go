package os

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// 运行程序的目录
	Getwd = os.Getwd
)

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
