// 更新操作
// https://gorm.io/docs/update.html
package gorm

import "gorm.io/gorm"

type Update struct {
	db *gorm.DB
}

// Update 更新数据表的一个字段
// model 为数据模型指针
//
// 必须指定查询条件
// 1. 如果设置了 model 中的主键，将作为查询条件，此时 query 可以为 nil
// 2. 如果没设置主键，则必须指定 query
func (u *Update) Update(
	model any,
	column string, value any,
	query any, args ...any,
) *gorm.DB {
	db := u.db.Model(model)
	if query != nil {
		db = db.Where(query, args...)
	}
	db = db.Update(column, value)
	return db
}
