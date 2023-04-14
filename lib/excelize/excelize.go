package excelize

import (
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

// GetRows 获取excel sheet所有行
func GetRows(filename, sheet string) ([][]string, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		logrus.Errorln("open file error", err)
		return nil, err
	}
	defer Close(f)

	rows, err := f.GetRows(sheet)
	return rows, err
}

// Close 关闭excel文件
func Close(f *excelize.File) {
	err := f.Close()
	if err != nil {
		logrus.Errorln("close file error", err)
	}
}
