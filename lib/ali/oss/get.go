// 下载文件
// https://help.aliyun.com/document_detail/32148.html
package oss

import (
	"io"
	"path"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

/*
流式下载
https://help.aliyun.com/document_detail/88617.htm
*/

// GetObject 读取文件字节
func GetObject(objectKey string, options ...oss.Option) ([]byte, error) {
	reader, err := Client.b.GetObject(objectKey, options...)
	if err != nil {
		logrus.Errorln("get object error", err)
		return nil, err
	}
	// 释放连接
	defer reader.Close()

	data, err := io.ReadAll(reader)
	return data, err
}

// GetObjectStr 读取文件字符串
func GetObjectStr(objectKey string, options ...oss.Option) (string, error) {
	bs, err := GetObject(objectKey, options...)
	if err != nil {
		logrus.Errorln("get object bytes error", err)
		return "", err
	}
	return string(bs), nil
}

// GetObjectToFile 下载文件 p指定目录 返回目录+文件名
func GetObjectToFile(objectKey, p string, options ...oss.Option) (string, error) {
	filename := path.Base(objectKey)
	file := filepath.Join(p, filename)
	err := Client.b.GetObjectToFile(objectKey, file, options...)
	return file, err
}

// Modify 修改文件，先读取文件字节，修改后上传
// options 是读取文件的选项
func Modify(f func([]byte) []byte, objectKey string, options ...oss.Option) error {
	bs, err := GetObject(objectKey, options...)
	if err != nil {
		logrus.Errorln("get object error", err)
		return err
	}
	bs = f(bs)
	err = PutObjectBytes(objectKey, bs)
	return err
}
