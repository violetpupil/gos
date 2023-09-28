package oss

import (
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	// 指定最大对象数
	MaxKeys = oss.MaxKeys
	// 指定分隔符 oss.Delimiter("/")
	Delimiter = oss.Delimiter
)

// Prefix 指定路径前缀
func Prefix(prefix string) oss.Option {
	// 不能以斜杆开头
	return oss.Prefix(strings.TrimPrefix(prefix, "/"))
}

// DelimiterSlash 指定目录前缀和/分隔符
// prefix传目录路径，如果没有后缀/，会自动添加
func DelimiterSlash(prefix string) []oss.Option {
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	return []oss.Option{Prefix(prefix), oss.Delimiter("/")}
}
