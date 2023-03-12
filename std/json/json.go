package json

import (
	"encoding/json"
	"os"

	"github.com/violetpupil/components/lib/logrus"
)

var (
	// json编码
	Marshal = json.Marshal
	// json解码
	Unmarshal = json.Unmarshal
)

// Load 加载json文件，v是非nil指针
func Load(name string, v any) error {
	bs, err := os.ReadFile(name)
	if err != nil {
		logrus.Error("read file error ", err)
		return err
	}
	err = json.Unmarshal(bs, v)
	return err
}

// Dump 写入json文件，并格式化
func Dump(v any, name string) error {
	// 缩进两个空格
	bs, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		logrus.Error("json marshal error ", err)
		return err
	}
	err = os.WriteFile(name, bs, 0666)
	return err
}
