// 管理文件
// https://help.aliyun.com/zh/oss/developer-reference/manage-objects-11
package oss

import (
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type (
	// oss对象
	// Key 对象名 目录以斜杆结尾
	ObjectProperties = oss.ObjectProperties
)

// IsDir 判断oss对象是否为目录
func IsDir(o oss.ObjectProperties) bool {
	return strings.HasSuffix(o.Key, "/")
}

// IsObjectExist 检查对象是否存在
func IsObjectExist(key string) (bool, error) {
	return Client.b.IsObjectExist(key)
}
