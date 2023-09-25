// 上传文件
// https://help.aliyun.com/document_detail/32147.html
package oss

import (
	"bytes"
	"io"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

/*
简单上传
https://help.aliyun.com/document_detail/88601.html
*/

// PutObjectFromFile 上传本地文件到oss
// objectKey 不能以斜杆开头
func PutObjectFromFile(objectKey, filePath string, options ...oss.Option) error {
	err := Client.b.PutObjectFromFile(objectKey, filePath, options...)
	return err
}

// PutObjectFromFileACL 上传本地文件到oss，并修改acl权限
func PutObjectFromFileACL(objectKey, filePath string, acl oss.ACLType) error {
	o := oss.ObjectACL(acl)
	err := PutObjectFromFile(objectKey, filePath, o)
	return err
}

// PutObject 放文件，如果文件存在会覆盖
func PutObject(objectKey string, reader io.Reader, options ...oss.Option) error {
	err := Client.b.PutObject(objectKey, reader, options...)
	return err
}

// PutObjectString 将字符串作为文件放到oss
func PutObjectString(objectKey string, s string, options ...oss.Option) error {
	err := PutObject(objectKey, strings.NewReader(s), options...)
	return err
}

// PutObjectBytes 将bytes作为文件放到oss
func PutObjectBytes(objectKey string, bs []byte, options ...oss.Option) error {
	err := PutObject(objectKey, bytes.NewReader(bs), options...)
	return err
}

// PutObjectURL 将url中的文件放到oss
func PutObjectURL(objectKey string, u string, options ...oss.Option) error {
	res, err := resty.New().R().Get(u)
	if err != nil {
		logrus.Errorln("get file error", err)
		return err
	}
	return PutObjectBytes(objectKey, res.Body(), options...)
}
