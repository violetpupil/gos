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
