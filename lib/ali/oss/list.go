// 列举文件
// https://help.aliyun.com/zh/oss/developer-reference/list-objects-10
package oss

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

type (
	// oss.Bucket.ListObjectsV2()结果
	// Objects 对象列表，有时包含指定目录对象
	// 指定oss.Delimiter("/")后, 所有后代目录及文件 -> 指定目录下文件
	// CommonPrefixes 以delimiter结尾的路径
	// 指定oss.Delimiter("/")后, 空切片 -> 指定目录下目录路径
	ListObjectsResultV2 = oss.ListObjectsResultV2
)

// ListObjectsAll 获取所有对象
// root设置为true，只获取指定目录下目录和文件，分别位于结果的CommonPrefixes和Objects字段
// root设置为false，获取所有后代目录及文件，位于结果的Objects字段
func ListObjectsAll(prefix string, root bool) (*oss.ListObjectsResultV2, error) {
	var out oss.ListObjectsResultV2
	// 翻页
	var token string
	options := make([]oss.Option, 0)
	if root {
		options = append(options, Root(prefix)...)
	} else {
		options = append(options, Prefix(prefix))
	}
	for {
		con := oss.ContinuationToken(token)
		options = append(options, con)
		res, err := Client.b.ListObjectsV2(options...)
		if err != nil {
			logrus.Errorln("list objects error", err)
			return nil, err
		}
		out.Objects = append(out.Objects, res.Objects...)
		out.CommonPrefixes = append(out.CommonPrefixes, res.CommonPrefixes...)

		if res.IsTruncated {
			token = res.NextContinuationToken
		} else {
			break
		}
	}
	return &out, nil
}

// ListObjects 获取指定目录下文件，max为数量
func ListObjects(prefix string, max int) ([]oss.ObjectProperties, error) {
	options := Root(prefix)
	options = append(options, oss.MaxKeys(max))
	res, err := Client.b.ListObjectsV2(options...)
	if err != nil {
		logrus.Errorln("list objects error", err)
		return nil, err
	}
	return res.Objects, nil
}

// ListObject 获取单个对象
func ListObject(key string) (*oss.ObjectProperties, error) {
	res, err := Client.b.ListObjectsV2(Prefix(key))
	if err != nil {
		logrus.Errorln("list objects error", err)
		return nil, err
	}
	if len(res.Objects) != 1 {
		logrus.Errorf("list objects length %d", len(res.Objects))
		return nil, fmt.Errorf("list objects length %d", len(res.Objects))
	}
	return &res.Objects[0], nil
}
