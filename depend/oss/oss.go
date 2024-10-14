package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
)

func Bucket(endpoint, accessKeyID, accessKeySecret, bucketName string) *oss.Bucket {
	c, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		zap.L().Error("new client error", zap.Error(err))
		return nil
	}
	b, err := c.Bucket(bucketName)
	if err != nil {
		zap.L().Error("get bucket error", zap.Error(err))
		return nil
	}
	return b
}
