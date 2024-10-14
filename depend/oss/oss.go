package oss

import (
	"io"
	"net/http"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
)

func Bucket(endpoint, accessKeyID, accessKeySecret, bucketName string) (*oss.Bucket, error) {
	c, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		zap.L().Error("new client error", zap.Error(err))
		return nil, err
	}
	b, err := c.Bucket(bucketName)
	if err != nil {
		zap.L().Error("get bucket error", zap.Error(err))
		return nil, err
	}
	return b, nil
}

func PutObjectFromURL(b *oss.Bucket, url string, key string) {
	log := zap.L().With(zap.String("key", key), zap.String("url", url))

	resp, err := http.Get(url)
	if err != nil {
		log.Error("get url error", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	err = b.PutObject(key, io.Reader(resp.Body))
	if err != nil {
		log.Error("put object error", zap.Error(err))
	} else {
		log.Info("put object success")
	}
}
