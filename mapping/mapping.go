package mapping

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
