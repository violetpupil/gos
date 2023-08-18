// 无需声明，直接获取json字段值
// fastjson parses the input JSON only once.
// Validates the parsed JSON.
// 数组使用0开始的索引
package fastjson

import "github.com/valyala/fastjson"

var (
	// 解析json字符串
	Parse = fastjson.Parse
	// 解析字段类型number，解析失败时，返回0
	GetInt    = fastjson.GetInt
	GetString = fastjson.GetString
)

type (
	// json值
	// Get() 获取json值，不存在返回nil
	// GetArray() 获取json值切片，不存在返回nil
	// GetBool() 获取bool
	// GetInt() 获取int
	// GetStringBytes() 获取bytes，不存在返回nil
	// Type() 值类型
	Value = fastjson.Value
)
