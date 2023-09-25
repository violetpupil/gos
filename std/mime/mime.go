// https://www.iana.org/assignments/media-types/media-types.xhtml
package mime

import "mime"

const (
	Json     = "application/json"
	FormUrl  = "application/x-www-form-urlencoded"
	FormData = "multipart/form-data"
	JPEG     = "image/jpeg"
)

var (
	// 根据mime类型返回文件扩展名列表
	// [.jfif .jpe .jpeg .jpg]
	// 当没有对应扩展名时，返回nil
	ExtensionsByType = mime.ExtensionsByType
)
