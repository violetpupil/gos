package url

import (
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

type (
	// URL url信息
	// RawQuery encoded query values, without '?'
	URL = url.URL
	// Values 查询字符串或表单，键大小写敏感
	Values = url.Values
)

var (
	// 查询字符串编码
	QueryEscape = url.QueryEscape
	// 查询字符串解码
	QueryUnescape = url.QueryUnescape
)

// Path 返回url的路径部分，无前缀斜杆
func Path(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		logrus.Errorln("url parse error", err)
		return "", err
	}
	// 参数是不以斜杆开头的路径时，u.Path 也没有前置斜杆
	// 其它情况 u.Path 会以斜杆开头
	path := strings.TrimPrefix(u.Path, "/")
	return path, nil
}
