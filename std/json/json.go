package json

import (
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"
)

var (
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
		zap.L().Error("json stdout error", zap.Error(err))
		fmt.Println(v)
	}
}
