// 权限控制
// https://help.aliyun.com/zh/oss/developer-reference/access-control-1/
package oss

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

// Access Control List权限
const (
	ACLPrivate         = oss.ACLPrivate         // 私有读，私有写，bucket默认
	ACLPublicRead      = oss.ACLPublicRead      // 公开读，私有写
	ACLPublicReadWrite = oss.ACLPublicReadWrite // 公开读，公开写
	ACLDefault         = oss.ACLDefault         // 适用于object，继承bucket权限，object默认
)

/*
授权访问
https://help.aliyun.com/zh/oss/developer-reference/authorize-access-5
*/

// SignURL 生成私有文件访问地址
// http://bucket.oss-cn-shenzhen.aliyuncs.com/tmp.txt?Expires=1695707083&OSSAccessKeyId=&Signature=
func SignURL(objectKey string, expiredInSec int64, options ...oss.Option) (string, error) {
	s, err := Client.b.SignURL(objectKey, oss.HTTPGet, expiredInSec, options...)
	return s, err
}
