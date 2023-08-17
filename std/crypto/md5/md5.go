package md5

import (
	"crypto/md5"
	"fmt"
)

// Sum md5摘要，结果为32位16进制字符串
func Sum(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

// SumSalt md5加盐摘要，结果为32位16进制字符串
func SumSalt(data, salt string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data+salt)))
}
