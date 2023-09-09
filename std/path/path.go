package path

import (
	"path"
	"strings"
)

// 函数
var (
	// Base 返回路径最后一部分
	// "/a" -> a
	// "/a/" -> a
	// "" -> .
	// "/" -> /
	Base = path.Base
	// Ext 返回文件扩展名 .txt
	Ext = path.Ext
	// Join 拼接路径，用/
	// 去掉结尾的/
	Join = path.Join
)

// Filename 获取路径中的文件名 /a.txt -> a
func Filename(p string) string {
	base := path.Base(p)
	ext := path.Ext(p)
	r := strings.TrimSuffix(base, ext)
	return r
}
