package gorm

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// crud 数据库操作
// 必须先调用 Init 初始化
type crud struct {
	db *gorm.DB
	C  *create
	R  *query
	U  *update
	D  *delete
}

var Crud *crud

func Init(db *gorm.DB) {
	Crud = new(crud)
	Crud.db = db
	Crud.C = &create{db}
	Crud.R = &query{db}
	Crud.U = &update{db}
	Crud.D = &delete{db}
}

// InitMySQL 初始化mysql数据库操作
// https://gorm.io/docs/connecting_to_the_database.html
// https://github.com/go-sql-driver/mysql
func InitMySQL(user, pass, host, port, d string) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, d,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Errorln("open mysql error", err)
		return err
	}
	Init(db)
	return nil
}

const (
	EnvHost     = "DBHost"
	EnvPort     = "DBPort"
	EnvUser     = "DBUser"
	EnvPass     = "DBPass"
	EnvDatabase = "DBDatabase"
)

// InitMySQLEnv 使用环境变量初始化mysql数据库操作
func InitMySQLEnv() error {
	host := os.Getenv(EnvHost)
	port := os.Getenv(EnvPort)
	user := os.Getenv(EnvUser)
	pass := os.Getenv(EnvPass)
	db := os.Getenv(EnvDatabase)
	err := InitMySQL(user, pass, host, port, db)
	return err
}

// AutoMigrate 自动迁移模型，不会删除列
// https://gorm.io/docs/migration.html
// dst是数据模型指针
func (c *crud) AutoMigrate(dst ...any) error {
	err := c.db.AutoMigrate(dst...)
	return err
}
