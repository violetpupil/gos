package gorm

import (
	"database/sql"

	"gorm.io/gorm"
)

type transaction struct {
	db *gorm.DB
}

// Transaction 在事务中调用fc
// fc返回错误则回滚，否则提交
func (t *transaction) Transaction(fc func(tx *DB) error, opts ...*sql.TxOptions) error {
	return t.db.Transaction(fc, opts...)
}
