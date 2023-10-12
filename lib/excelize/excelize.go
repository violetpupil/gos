package excelize

import (
	"errors"
	"fmt"
	"strings"

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

// NewFile 创建excel文件 file是文件路径 文件名不带后缀或带.xlsx
// 最多写26列
func NewFile(rows [][]any, file string) error {
	f := excelize.NewFile()
	defer Close(f)

	for i, row := range rows {
		if len(row) > 26 {
			return errors.New("max 26 column")
		}
		for j, cell := range row {
			// A的unicode码点是65
			f.SetCellValue("Sheet1", fmt.Sprintf("%c%d", 65+j, i+1), cell)
		}
	}
	if !strings.HasSuffix(file, ".xlsx") {
		file += ".xlsx"
	}
	err := f.SaveAs(file)
	return err
}
