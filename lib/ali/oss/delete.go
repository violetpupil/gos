// https://help.aliyun.com/zh/oss/developer-reference/delete-objects-11
package oss

import (
	"errors"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

// DeleteObjects 删除多个对象
// 文件夹以/结尾
func DeleteObjects(keys ...string) error {
	// 不返回删除结果
	_, err := Client.b.DeleteObjects(keys, oss.DeleteObjectsQuiet(true))
	return err
}

// DeleteAll 删除所有指定前缀对象，f返回true表示忽略
// 设置test表示只输出删除对象，不实际删除
func DeleteAll(prefix string, f func(oss.ObjectProperties) bool, test bool) error {
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
				if IsDir(o) {
					logrus.Infoln("try delete dir", o.Key)
				} else {
					logrus.Infoln("will delete object", o.Key)
				}
			} else {
				logrus.Infoln("skip object", o.Key)
			}
		}
		if !test && len(keys) > 0 {
			_, err = Client.b.DeleteObjects(keys, oss.DeleteObjectsQuiet(true))
			if err != nil {
				logrus.Errorln("delete objects error", err)
				return err
			}
			logrus.Infoln("delete objects success")
		}

		if res.IsTruncated {
			token = res.NextContinuationToken
		} else {
			break
		}
	}
	return nil
}
