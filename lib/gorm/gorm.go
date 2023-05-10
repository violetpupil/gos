// https://gorm.io/
// https://gorm.io/docs/
package gorm

import (
	"fmt"

	"github.com/caarlos0/env/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// crud 数据库操作
// 必须先调用 Init 初始化
type crud struct {
	db  *gorm.DB
	C   *create
	R   *query
	U   *update
	D   *delete
	Raw *raw
}

var Crud *crud

func Init(db *gorm.DB) {
	Crud = new(crud)
	Crud.db = db
	Crud.C = &create{db}
	Crud.R = &query{db}
	Crud.U = &update{db}
	Crud.D = &delete{db}
	Crud.Raw = &raw{db}
}

// Config mysql配置
type Config struct {
	Host     string `env:"MySQLHost"`
	Port     string `env:"MySQLPort"`
	User     string `env:"MySQLUser"`
	Pass     string `env:"MySQLPass"`
	Database string `env:"MySQLDatabase"`
	// gorm.Config
	Logger logger.Interface
	// 迁移模型时，不创建外键约束
	DisableForeignKeyConstraintWhenMigrating bool
}

// InitMySQL 初始化mysql数据库操作
// https://gorm.io/docs/connecting_to_the_database.html
// https://github.com/go-sql-driver/mysql
func InitMySQL(c Config) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Pass, c.Host, c.Port, c.Database,
	)
	dialector := mysql.Open(dsn)

	opt := &gorm.Config{
		Logger:                                   c.Logger,
		DisableForeignKeyConstraintWhenMigrating: c.DisableForeignKeyConstraintWhenMigrating,
	}
	db, err := gorm.Open(dialector, opt)
	if err != nil {
		logrus.Errorln("open mysql error", err)
		return err
	}
	Init(db)
	return nil
}

// InitMySQLEnv 使用环境变量初始化mysql数据库操作
func InitMySQLEnv() error {
	var c Config
	err := env.Parse(&c)
	if err != nil {
		logrus.Errorln("env parse error", err)
		return err
	}
	err = InitMySQL(c)
	return err
}

// AutoMigrate 自动迁移模型，不会删除列
// https://gorm.io/docs/migration.html
// opt是创建表选项 COMMENT='表注释'
// dst是数据模型指针
// 默认创建数据库外键约束
func (c *crud) AutoMigrate(opt string, dst ...any) error {
	if opt != "" {
		return c.db.Set("gorm:table_options", opt).AutoMigrate(dst...)
	} else {
		return c.db.AutoMigrate(dst...)
	}
}
