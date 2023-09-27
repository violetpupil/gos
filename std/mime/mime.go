// https://www.iana.org/assignments/media-types/media-types.xhtml
package mime

import (
	"mime"

	"github.com/sirupsen/logrus"
)

const (
	Json     = "application/json"
	FormUrl  = "application/x-www-form-urlencoded"
	FormData = "multipart/form-data"
	JPEG     = "image/jpeg"
	Excel    = "application/vnd.ms-excel"
)

// ExtensionByType 根据mime类型返回文件扩展名
// 当没有对应扩展名时，返回空字符串
func ExtensionByType(typ string) (string, error) {
	// 根据mime类型返回文件扩展名列表
	// [.jfif .jpe .jpeg .jpg]
	// 当没有对应扩展名时，返回nil
	exts, err := mime.ExtensionsByType(typ)
	if err != nil {
		logrus.Errorln("extensions by type error", err)
		return "", err
	}
	if len(exts) == 0 {
		return "", nil
	} else {
		return exts[len(exts)-1], nil
	}
}
