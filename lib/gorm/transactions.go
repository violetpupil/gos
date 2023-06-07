// https://gorm.io/docs/transactions.html
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
// 注意fc中使用参数tx执行操作
// fc返回的错误会传递出来
func (t *transaction) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return t.db.Transaction(fc, opts...)
}
