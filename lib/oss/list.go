// 列举文件
// https://help.aliyun.com/document_detail/88639.html
package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

// ListObjects 获取对象列表，prefix不能以斜杆开头
func ListObjects(prefix string) ([]oss.ObjectProperties, error) {
	objects := make([]oss.ObjectProperties, 0)
	pre := oss.Prefix(prefix)
	token := ""
	for {
		con := oss.ContinuationToken(token)
		res, err := Client.b.ListObjectsV2(con, pre)
		if err != nil {
			logrus.Errorln("list objects error", err)
			return nil, err
		}
		objects = append(objects, res.Objects...)

		// 默认一次返回100条记录
		if res.IsTruncated {
			token = res.NextContinuationToken
		} else {
			break
		}
	}
	return objects, nil
}
