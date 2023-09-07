// https://github.com/aliyun/oss-browser
package oss

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/sirupsen/logrus"
)

// AuthToken 授权码结构
// https://github.com/aliyun/oss-browser/blob/develop/docs/authToken.md
type AuthToken struct {
	ID         string `json:"id"`         // access key id
	Secret     string `json:"secret"`     // access key secret
	SToken     string `json:"stoken"`     // security token
	Privilege  string `json:"privilege"`  // 权限
	Expiration string `json:"expiration"` // 过期时间
	OssPath    string `json:"osspath"`    // oss路径 oss://bucket/path
	Region     string `json:"region"`     // 区域
	EpTpl      string `json:"eptpl"`      // endpoint
	Bucket     string `json:"bucket"`     // 桶，从OssPath解析得到
}

// ParseAuthToken 解析授权码
func ParseAuthToken(token string) (*AuthToken, error) {
	var t AuthToken
	bs, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		logrus.Errorln("base64 decode error", err)
		return nil, err
	}
	err = json.Unmarshal(bs, &t)
	if err != nil {
		logrus.Errorln("json unmarshal error", err)
		return nil, err
	}

	path := strings.Split(t.OssPath, "/")
	if len(path) >= 3 {
		t.Bucket = path[2]
	}
	return &t, nil
}
