// https://gorm.io/docs/security.html
package gorm

import (
	"strings"

	"golang.org/x/exp/slices"
)

// OrderCheck 排序字段检查
func OrderCheck(fields []string, s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, ",", " ")
	parts := strings.Split(s, " ")
	for _, p := range parts {
		if p == "" || p == "asc" || p == "desc" {
			continue
		}
		if slices.Contains(fields, p) {
			continue
		}
		return false
	}
	return true
}
