package json

import (
	"encoding/json"
)

var (
	// 字段名大小写不敏感
	// 结构体使用指针字段，可以区分不存在还是赋了默认值
	//
	// 当json解码到any时，接口实际类型对应关系
	// []interface{}, for JSON arrays
	// map[string]interface{}, for JSON objects
	Unmarshal = json.Unmarshal
)
