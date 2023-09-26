package url

import (
	"net/url"
	"path"
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

// HasExt 路径中是否有文件扩展名
func HasExt(rawURL string) (bool, error) {
	p, err := Path(rawURL)
	if err != nil {
		logrus.Errorln("parse path error", err)
		return false, err
	}
	return path.Ext(p) != "", nil
}

// Query 获取查询字符串字段单个值，不存在返回空
func Query(rawURL, key string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		logrus.Errorln("url parse error", err)
		return "", err
	}
	values, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		logrus.Errorln("parse query error", err)
		return "", err
	}
	return values.Get(key), nil
}
