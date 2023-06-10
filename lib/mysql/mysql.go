package mysql

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

const (
	MySQLErrorNumberDuplicateEntry = 1062
)

type (
	// mysql返回的错误
	MySQLError = mysql.MySQLError
)

// AsMySQLError 查找错误链中的mysql.MySQLError并返回
func AsMySQLError(err error) *mysql.MySQLError {
	e := new(mysql.MySQLError)
	ok := errors.As(err, &e)
	if !ok {
		return nil
	}
	return e
}

// IsDuplicateEntry 判断错误是否为mysql重复实体
func IsDuplicateEntry(err error) bool {
	e := AsMySQLError(err)
	if e == nil {
		return false
	}
	return e.Number == MySQLErrorNumberDuplicateEntry
}
