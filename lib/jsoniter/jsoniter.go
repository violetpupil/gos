// 100% compatibility with standard lib
package jsoniter

import jsoniter "github.com/json-iterator/go"

var (
	// json 编码，返回字符串
	MarshalToString = jsoniter.MarshalToString
)
