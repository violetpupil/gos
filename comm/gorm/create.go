// 插入操作
// https://gorm.io/docs/create.html
package gorm

import "gorm.io/gorm"

type create struct {
	db *gorm.DB
}

// Create 插入数据，返回插入记录数
// value为数据模型指针或指针切片
// 如果主键是插入时生成的，会自动更新
func (c *create) Create(value any) (int64, error) {
	db := c.db.Create(value)
	return db.RowsAffected, db.Error
}
