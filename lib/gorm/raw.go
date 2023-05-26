// https://gorm.io/docs/sql_builder.html
package gorm

import "gorm.io/gorm"

type raw struct {
	db *gorm.DB
}

// Exec 直接执行sql语句，不能获取查询结果
func (r *raw) Exec(sql string, values ...interface{}) error {
	err := r.db.Exec(sql, values...).Error
	return err
}

// RawScan 直接执行sql语句，并获取结果
func (r *raw) RawScan(sql string, dest interface{}, values ...interface{}) error {
	err := r.db.Raw(sql, values...).Scan(dest).Error
	return err
}
