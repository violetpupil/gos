// 阿里云
// 语音处理 https://help.aliyun.com/product/30413.html
package ali

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/lib/gocsv"
)

// Lang 语种
// 机器翻译 https://help.aliyun.com/product/30396.html
// 使用python获取语言代码列表
// pandas.read_html("https://help.aliyun.com/document_detail/330911.html")[0].to_csv("lang.csv", encoding='utf8')
type Lang struct {
	No     int    `csv:"序号"`
	NameZh string `csv:"语种中文名称"`
	NameEn string `csv:"语种英文名称"`
	Code   string `csv:"语言代码"`
}

// WriteLangs 将语种csv转为go代码
func WriteLangs() {
	var ls []Lang
	err := gocsv.UnmarshalFile("./lang.csv", &ls)
	if err != nil {
		logrus.Error("unmarshal file error ", err)
		return
	}

	s := fmt.Sprintf("%#v", ls)
	s = strings.ReplaceAll(s, "ali.", "")
	s = strings.ReplaceAll(s, "Lang{No", "{No")
	s = "package ali\n\nvar Langs=" + s
	err = os.WriteFile("lang.go", []byte(s), 0666)
	logrus.WithField("Error", err).Info("write file")
}
