package misc

import (
	"path"
	"strings"
)

// Filename 获取路径中的文件名 /a.txt -> a
func Filename(p string) string {
	base := path.Base(p)
	ext := path.Ext(p)
	r := strings.TrimSuffix(base, ext)
	return r
}
