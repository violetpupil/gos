package depend

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
)

func NewBucket(log *zap.Logger, endpoint, accessKeyID, accessKeySecret, bucketName string) (*oss.Bucket, error) {
	c, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		log.Error("new client error", zap.Error(err))
		return nil, err
	}
	b, err := c.Bucket(bucketName)
	if err != nil {
		log.Error("get bucket error", zap.Error(err))
		return nil, err
	}
	return b, nil
}
