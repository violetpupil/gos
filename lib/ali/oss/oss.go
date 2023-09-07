// oss客户端，必须先调用 Init 初始化
// https://help.aliyun.com/zh/oss/
// https://help.aliyun.com/zh/oss/developer-reference/introduction-3
package oss

import (
	"errors"
	"os"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

type client struct {
	c *oss.Client
	b *oss.Bucket
}

var Client *client

// Init 初始化oss客户端
// endpoint 是oss数据中心地址 例如https://oss-cn-hangzhou.aliyuncs.com
// https://help.aliyun.com/zh/oss/user-guide/regions-and-endpoints
func Init(
	endpoint, accessKeyID, accessKeySecret, bucketName string,
	options ...oss.ClientOption,
) error {
	c, err := oss.New(endpoint, accessKeyID, accessKeySecret, options...)
	if err != nil {
		logrus.Error("new client error ", err)
		return err
	}
	b, err := c.Bucket(bucketName)
	if err != nil {
		logrus.Error("get bucket error ", err)
		return err
	}
	Client = &client{c: c, b: b}
	return err
}

// InitToken 初始化oss客户端，使用oss browser auth token
// 返回token可访问的路径前缀 path/
// https://github.com/aliyun/oss-browser/blob/develop/docs/authToken.md
// https://help.aliyun.com/zh/oss/developer-reference/go-configure-access-credentials
func InitToken(token string) (string, error) {
	t, err := ParseAuthToken(token)
	if err != nil {
		logrus.Errorln("parse auth token error", err)
		return "", err
	}
	if t.ExpireTs != 0 && time.Now().Unix() > t.ExpireTs {
		logrus.Infoln("auth token expired")
		return "", errors.New("auth token expired")
	}

	err = Init(t.EpTpl, t.ID, t.Secret, t.Bucket, oss.SecurityToken(t.SToken))
	if err != nil {
		logrus.Infoln("init error", err)
		return "", err
	}
	return t.PathPrefix, nil
}

// InitEnv 使用环境变量初始化oss客户端
// https://help.aliyun.com/zh/oss/developer-reference/go-configure-access-credentials
func InitEnv() error {
	var options []oss.ClientOption
	token := os.Getenv("OSS_SESSION_TOKEN")
	if token != "" {
		options = append(options, oss.SecurityToken(token))
	}
	err := Init(
		os.Getenv("OSS_ENDPOINT"),
		os.Getenv("OSS_ACCESS_KEY_ID"),
		os.Getenv("OSS_ACCESS_KEY_SECRET"),
		os.Getenv("OSS_BUCKET"),
		options...,
	)
	return err
}
