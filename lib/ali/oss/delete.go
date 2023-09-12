// https://help.aliyun.com/zh/oss/developer-reference/delete-objects-11
package oss

import (
	"errors"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

// DeleteObject 删除单个对象，只能删除空文件夹
func DeleteObject(key string) error {
	err := Client.b.DeleteObject(key)
	return err
}

// DeleteObjects 删除多个对象
func DeleteObjects(keys ...string) error {
	// 不返回删除结果
	_, err := Client.b.DeleteObjects(keys, oss.DeleteObjectsQuiet(true))
	return err
}

// DeleteAll 删除所有指定前缀对象，f返回true表示忽略
func DeleteAll(prefix string, f func(oss.ObjectProperties) bool) error {
	// 禁止清空桶
	if prefix == "" {
		return errors.New("prefix is empty")
	}

	var token string
	for {
		res, err := Client.b.ListObjectsV2(Prefix(prefix), oss.ContinuationToken(token))
		if err != nil {
			logrus.Errorln("list objects error", err)
			return err
		}
		if len(res.Objects) == 0 {
			break
		}

		var keys []string
		for _, o := range res.Objects {
			if !f(o) {
				keys = append(keys, o.Key)
			}
		}
		_, err = Client.b.DeleteObjects(keys, oss.DeleteObjectsQuiet(true))
		if err != nil {
			logrus.Errorln("delete objects error", err)
			return err
		}

		if res.IsTruncated {
			token = res.NextContinuationToken
		} else {
			break
		}
	}
	return nil
}
