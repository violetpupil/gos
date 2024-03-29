package strings

import "strings"

var (
	// Split 用分隔符获取子串
	// 如果分隔符不是空，字符串中不包含分隔符，结果是原字符串切片
	Split = strings.Split
	// SplitN 用分隔符获取子串，从左限定结果数量，数量<0相当于Split
	SplitN = strings.SplitN
	// TrimSuffix 去除后缀
	TrimSuffix = strings.TrimSuffix
)

type (
	// 字符串构造器
	// Write() 写字节数组
	// WriteByte() 写单个字节
	// WriteString() 写字符串
	Builder = strings.Builder
)

// SplitLast 用分隔符获取最后一个子串
// 源字符串结尾的分隔符会先去掉
func SplitLast(s, sep string) string {
	s = strings.TrimSuffix(s, sep)
	sub := strings.Split(s, sep)
	return sub[len(sub)-1]
}
