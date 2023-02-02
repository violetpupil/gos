package gorm

import (
	"errors"

	"gorm.io/gorm"
)

// RecordNotFound 检查查询函数错误是否为 ErrRecordNotFound
// First, Take, Last 会返回这个错误
// https://gorm.io/docs/query.html
func RecordNotFound(e error) bool {
	return errors.Is(e, gorm.ErrRecordNotFound)
}
