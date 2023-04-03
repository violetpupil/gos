// map不初始化，只能读，不能写
// 使用不存在的键取值，会返回默认值
package types

import (
	"fmt"
	"strings"
)

// Join 将映射拼接成字符串，映射键值都是字符串类型
func Join(m map[string]string, kvSep, lineSep string) string {
	var s []string
	for k, v := range m {
		s = append(s, fmt.Sprint(k, kvSep, v))
	}
	return strings.Join(s, lineSep)
}

// Swap 交换map键值，如果值有重复的，会被覆盖
func Swap(s map[string]string) map[string]string {
	d := make(map[string]string)
	for k, v := range s {
		d[v] = k
	}
	return d
}
