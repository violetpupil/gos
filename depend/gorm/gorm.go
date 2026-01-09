package gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func NewMysqlDialector(c Config) gorm.Dialector {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Pass, c.Host, c.Port, c.Database,
	)
	config := mysql.Config{
		DSN: dsn,
	}
	dialector := mysql.New(config)
	return dialector
}

func NewMysqlDB(c Config) (*gorm.DB, error) {
	dialector := NewMysqlDialector(c)

	opt := &gorm.Config{}
	db, err := gorm.Open(dialector, opt)
	return db, err
}
