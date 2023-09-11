package oss

import (
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	// 指定最大对象数
	MaxKeys = oss.MaxKeys
)

// Prefix 指定路径前缀
func Prefix(prefix string) oss.Option {
	// 不能以斜杆开头
	return oss.Prefix(strings.TrimPrefix(prefix, "/"))
}

// Root 只获取指定目录下目录及文件
func Root(prefix string) []oss.Option {
	// prefix必须斜杆结尾
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	// 指定oss.Delimiter("/")后
	// oss.ListObjectsResultV2.Objects
	// 所有后代目录及文件 -> 指定目录下文件
	// oss.ListObjectsResultV2.CommonPrefixes
	// 空切片 -> 指定目录下目录路径
	return []oss.Option{Prefix(prefix), oss.Delimiter("/")}
}
