package yaml

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// UnmarshalFile 读取 yaml 文件并加载到结构体指针
// 结构体字段使用 yaml 标签指定键名
// 如果指定了键名，则大小写敏感匹配
// 否则使用字段名，大小写不敏感匹配
func UnmarshalFile(f string, v any) error {
	bs, err := os.ReadFile(f)
	if err != nil {
		logrus.Errorln("read file error", err)
		return err
	}
	err = yaml.Unmarshal(bs, v)
	return err
}
