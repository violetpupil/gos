// csv文件处理
// 字段可以用""包裹，""内可以换行，也可以使用""表示单个"
package csv

import (
	"encoding/csv"
	"os"

	"github.com/violetpupil/components/lib/logrus"
)

// ReadAll 读取csv文件
func ReadAll(name string) (records [][]string, err error) {
	f, err := os.Open(name)
	if err != nil {
		logrus.Error("open error ", err)
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err = reader.ReadAll()
	return
}
