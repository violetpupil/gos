// 下载文件
// https://help.aliyun.com/document_detail/32148.html
package oss

import (
	"io"
	"os"
	"path"

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

// Download 下载文件
func Download(objectKey string, options ...oss.Option) error {
	reader, err := Client.b.GetObject(objectKey, options...)
	if err != nil {
		logrus.Errorln("get object error", err)
		return err
	}
	// 释放连接
	defer reader.Close()

	// 创建同名文件
	filename := path.Base(objectKey)
	fd, err := os.Create(filename)
	if err != nil {
		logrus.Errorln("create file error", err)
		return err
	}
	defer fd.Close()

	_, err = io.Copy(fd, reader)
	return err
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
