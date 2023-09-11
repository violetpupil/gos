// 列举文件
// https://help.aliyun.com/zh/oss/developer-reference/list-objects-10
package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

// ListObjects 获取对象列表
// prefix不能以斜杆开头
// 包括目录，目录以斜杆结尾
// first表示是否只获取第一页
func ListObjects(prefix string, first bool) ([]oss.ObjectProperties, error) {
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
		if !first && res.IsTruncated {
			token = res.NextContinuationToken
		} else {
			break
		}
	}
	return objects, nil
}

// ListObjectsMaxKeys 获取对象列表 指定个数
// prefix不能以斜杆开头
// 包括目录，目录以斜杆结尾
func ListObjectsMaxKeys(prefix string, max int) ([]oss.ObjectProperties, error) {
	res, err := Client.b.ListObjectsV2(oss.Prefix(prefix), oss.MaxKeys(max))
	if err != nil {
		logrus.Errorln("list objects error", err)
		return nil, err
	}
	return res.Objects, nil
}

// ListObjectsDir 获取子目录列表
// prefix不能以斜杆开头，应该以斜杆结尾
func ListObjectsDir(prefix string) ([]string, error) {
	dirs := make([]string, 0)
	pre := oss.Prefix(prefix)
	token := ""
	for {
		con := oss.ContinuationToken(token)
		// 设置了oss.Delimiter("/")
		// res.Objects只有当前目录下文件，没有子目录及子目录下文件
		res, err := Client.b.ListObjectsV2(con, pre, oss.Delimiter("/"))
		if err != nil {
			logrus.Errorln("list objects error", err)
			return nil, err
		}
		dirs = append(dirs, res.CommonPrefixes...)

		if res.IsTruncated {
			token = res.NextContinuationToken
		} else {
			break
		}
	}
	return dirs, nil
}
