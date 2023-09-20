package json

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// json编码
	// 结构体标签 `json:","`
	// omitempty选项 当字段为默认值时，不编码
	// string选项 将数字、布尔编码成字符串 "true"
	Marshal = json.Marshal
	// json解码
	// 字段名大小写不敏感
	// 结构体使用指针字段，可以区分不存在还是赋了默认值
	//
	// 当json解码到any时，接口实际类型对应关系
	// bool, for JSON booleans
	// float64, for JSON numbers
	// string, for JSON strings
	// []interface{}, for JSON arrays
	// map[string]interface{}, for JSON objects
	// nil for JSON null
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

// MarshalIndent json编码并缩进
func MarshalIndent(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

// Stdout json编码并缩进，打印在标准输出并换行
// 编码失败，回退用 fmt.Println 打印
func Stdout(v any) {
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "  ")
	// 将编码结果写入writer，并添加换行
	err := e.Encode(v)
	if err != nil {
		logrus.Errorln("json stdout error", err)
		fmt.Println(v)
	}
}
