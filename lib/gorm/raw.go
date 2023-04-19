// https://gorm.io/docs/sql_builder.html
package gorm

import "gorm.io/gorm"

type raw struct {
	db *gorm.DB
}

// Exec 直接执行sql语句
func (r *raw) Exec(sql string, values ...interface{}) error {
	err := r.db.Exec(sql, values...).Error
	return err
}
