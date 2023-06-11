// 当json解码到any接口时，接口实际类型对应关系
// map[string]interface{}, for JSON objects
// https://pkg.go.dev/encoding/json#Unmarshal
//
// mapstructure可以直接将map[string]interface{}再解码
// When decoding to a struct, mapstructure will use the field name by default to perform the mapping.
// 大小写不敏感
// The default struct tag that mapstructure looks for is "mapstructure"
package mapstructure

import "github.com/goinggo/mapstructure"

var (
	// 将map解码为结构体
	Decode = mapstructure.Decode
)
