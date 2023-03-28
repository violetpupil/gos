// oss客户端，必须先调用 Init 初始化
// https://help.aliyun.com/product/31815.html
// https://help.aliyun.com/document_detail/32144.html
package oss

import (
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/violetpupil/components/lib/logrus"
)

type client struct {
	c *oss.Client
	b *oss.Bucket
}

var Client *client

// Init 初始化oss客户端
// endpoint 是oss数据中心地址 例如https://oss-cn-hangzhou.aliyuncs.com
// https://help.aliyun.com/document_detail/31837.html
func Init(endpoint, accessKeyID, accessKeySecret, bucketName string) error {
	c, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		logrus.Error("new client error ", err)
		return err
	}
	b, err := c.Bucket(bucketName)
	if err != nil {
		logrus.Error("get bucket error ", err)
		return err
	}
	Client = &client{c, b}
	return err
}

// InitEnv 使用环境变量初始化oss客户端
func InitEnv() {
	Init(
		os.Getenv("OssEndpoint"),
		os.Getenv("OssKeyId"),
		os.Getenv("OssKeySecret"),
		os.Getenv("OssBucket"),
	)
}

// PutObjectFromFile 上传本地文件到oss
// objectKey 不需要以斜杆开头
func PutObjectFromFile(objectKey, filePath string) error {
	err := Client.b.PutObjectFromFile(objectKey, filePath)
	return err
}
