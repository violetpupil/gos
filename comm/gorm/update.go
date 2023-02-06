// 更新操作
// https://gorm.io/docs/update.html
package gorm

import "gorm.io/gorm"

type update struct {
	db *gorm.DB
}

// Update 更新数据表的一个字段，返回更新记录条数
// model 为数据模型指针
//
// 必须指定查询条件
// 1. 如果设置了 model 中的主键，将作为查询条件，此时 query 可以为 nil
// 2. 如果没设置主键，则必须指定 query
func (u *update) Update(
	model any,
	column string, value any,
	query any, args ...any,
) (int64, error) {
	db := u.db.Model(model)
	if query != nil {
		db = db.Where(query, args...)
	}
	db = db.Update(column, value)
	return db.RowsAffected, db.Error
}

// Updates 更新数据表的多个字段，返回更新记录条数
// model 为数据模型指针
//
// values 必须为结构体或者 map[string]any
// values 为结构体时，只会更新非0字段
// 使用 Select, Omit 可以指定更新的字段
//
// 必须指定查询条件
// 1. 如果设置了 model 中的主键，将作为查询条件，此时 query 可以为 nil
// 2. 如果没设置主键，则必须指定 query
func (u *update) Updates(
	model any,
	values any,
	query any, args ...any,
) (int64, error) {
	db := u.db.Model(model)
	if query != nil {
		db = db.Where(query, args...)
	}
	db = db.Updates(values)
	return db.RowsAffected, db.Error
}
