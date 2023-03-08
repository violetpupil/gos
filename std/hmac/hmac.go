package hmac

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"hash"
)

// Sum hmac消息认证码
// h 哈希生成函数，例如 sha1.New sha256.New
// 结果为变长16进制字符串
func Sum(h func() hash.Hash, key string, s string) string {
	return fmt.Sprintf("%x", hmac.New(h, []byte(key)).Sum([]byte(s)))
}

// HmacSha1 hmac消息认证码sha1哈希摘要，结果为变长16进制字符串
func HmacSha1(key, s string) string {
	return Sum(sha1.New, key, s)
}
