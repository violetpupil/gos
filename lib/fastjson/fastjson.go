// 无需声明，直接获取json字段值
// https://pkg.go.dev/github.com/valyala/fastjson
// fastjson parses the input JSON only once.
// Validates the parsed JSON.
package fastjson

import "github.com/valyala/fastjson"

var (
	// 获取int字段值，解析失败时，返回0
	GetInt    = fastjson.GetInt
	GetString = fastjson.GetString
)
