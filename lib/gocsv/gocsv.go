package gocsv

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/violetpupil/components/lib/logrus"
)

// UnmarshalFile 读取csv文件
// 可以使用结构体标签csv，与文件第一行对应
func UnmarshalFile(name string, out any) error {
	f, err := os.Open(name)
	if err != nil {
		logrus.Error("open error ", err)
		return err
	}
	defer f.Close()

	err = gocsv.UnmarshalFile(f, out)
	return err
}
