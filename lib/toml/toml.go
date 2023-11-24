package toml

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/sirupsen/logrus"
)

// UnmarshalFile 读取 toml 文件并加载到结构体指针
// 结构体字段使用 toml 标签指定键名
// 不指定，则使用字段名
// 指定与否，都是大小写不敏感匹配
func UnmarshalFile(f string, v any) error {
	bs, err := os.ReadFile(f)
	if err != nil {
		logrus.Error("read file error ", err)
		return err
	}
	err = toml.Unmarshal(bs, v)
	return err
}
