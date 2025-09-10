package hmac

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"hash"
)

// Sum hmac消息认证码
// h 哈希生成函数，例如 sha1.New sha256.New
func Sum(h func() hash.Hash, key string, s string) []byte {
	mac := hmac.New(h, []byte(key))
	mac.Write([]byte(s))
	return mac.Sum(nil)
}

// HmacSha1 hmac消息认证码sha1哈希摘要
func HmacSha1(key string, s string) []byte {
	return Sum(sha1.New, key, s)
}

// HmacSha1Hex hmac消息认证码sha1 16进制摘要
func HmacSha1Hex(key, s string) string {
	return fmt.Sprintf("%x", Sum(sha1.New, key, s))
}

// HmacSha1Base64 hmac消息认证码sha1 base64摘要
func HmacSha1Base64(key, s string) string {
	return base64.StdEncoding.EncodeToString(HmacSha1(key, s))
}
