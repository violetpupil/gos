// 权限控制
// https://help.aliyun.com/document_detail/31833.html
package oss

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

// Access Control List权限
const (
	ACLPrivate         = oss.ACLPrivate         // 私有读，私有写，bucket默认
	ACLPublicRead      = oss.ACLPublicRead      // 公开读，私有写
	ACLPublicReadWrite = oss.ACLPublicReadWrite // 公开读，公开写
	ACLDefault         = oss.ACLDefault         // 适用于object，继承bucket权限，object默认
)
