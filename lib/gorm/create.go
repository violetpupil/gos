// 插入操作
// https://gorm.io/docs/create.html
package gorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type create struct {
	db *gorm.DB
}

// Create 插入数据，返回插入记录数
// value为数据模型指针或指针切片
// 如果主键是插入时生成的，会自动更新
// 数据被唯一性约束时，会报错
func (c *create) Create(value any) (int64, error) {
	db := c.db.Create(value)
	return db.RowsAffected, db.Error
}

// CreateDoNothing 插入数据，返回插入记录数
// value为数据模型指针或指针切片
// 如果主键是插入时生成的，会自动更新
// 数据被唯一性约束时，忽略该数据
func (c *create) CreateDoNothing(value any) (int64, error) {
	db := c.db.Clauses(clause.OnConflict{DoNothing: true}).Create(value)
	return db.RowsAffected, db.Error
}
