package gocsv

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/sirupsen/logrus"
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

// SetCSVReader 改变全局reader的字段分隔符，默认是逗号,
func SetCSVReader(comma rune) {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = comma
		return r
	})
}
