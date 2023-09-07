// oss客户端，必须先调用 Init 初始化
// https://help.aliyun.com/zh/oss/
// https://help.aliyun.com/zh/oss/developer-reference/introduction-3
//
// oss browser https://github.com/aliyun/oss-browser
// 授权码 https://github.com/aliyun/oss-browser/blob/develop/docs/authToken.md
package oss

import (
	"os"

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

// InitToken 初始化oss客户端，使用sts临时安全token
func InitToken(
	endpoint, accessKeyID, accessKeySecret, bucketName string,
	token string,
) error {
	err := Init(
		endpoint, accessKeyID, accessKeySecret, bucketName,
		oss.SecurityToken(token),
	)
	return err
}

// InitEnv 使用环境变量初始化oss客户端
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
