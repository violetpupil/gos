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
	// Objects 对象列表 指定prefix/前缀时，有时会包含prefix/对象
	// CommonPrefixes 被oss.Delimiter("/")分隔的目录路径
	ListObjectsResultV2 = oss.ListObjectsResultV2
)

// First 获取指定前缀的第一个对象
func First(prefix string) (*oss.ObjectProperties, error) {
	res, err := Client.b.ListObjectsV2(Prefix(prefix), oss.MaxKeys(1))
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

// ListObjects 获取多个对象
func ListObjects(options ...oss.Option) (*oss.ListObjectsResultV2, error) {
	result := new(oss.ListObjectsResultV2)
	var contToken string
	for {
		var os []oss.Option
		os = append(os, options...)
		os = append(os, oss.ContinuationToken(contToken))
		res, err := Client.b.ListObjectsV2(os...)
		if err != nil {
			logrus.Errorln("list objects error", err)
			return nil, err
		}
		result.Objects = append(result.Objects, res.Objects...)
		result.CommonPrefixes = append(result.CommonPrefixes, res.CommonPrefixes...)

		if res.IsTruncated {
			contToken = res.NextContinuationToken
		} else {
			break
		}
	}
	return result, nil
}

// ListObjectsPrefix 获取多个对象 指定前缀
// 自动去掉prefix开头的/
func ListObjectsPrefix(prefix string) (*oss.ListObjectsResultV2, error) {
	return ListObjects(Prefix(prefix))
}

// ListObjectsSlash 获取多个对象 指定目录前缀和/分隔符
// prefix传目录路径，如果没有后缀/，会自动添加
// 自动去掉prefix开头的/
//
// 返回
// oss.ListObjectsResultV2.Objects 指定目录下文件 有时会包含指定的目录前缀
// oss.ListObjectsResultV2.CommonPrefixes 指定目录下子目录
func ListObjectsSlash(prefix string) (*oss.ListObjectsResultV2, error) {
	return ListObjects(DelimiterSlash(prefix)...)
}
