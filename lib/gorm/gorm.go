// https://gorm.io/
// https://gorm.io/docs/
package gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// crud 数据库操作
type crud struct {
	db  *gorm.DB
	C   *create
	R   *query
	U   *update
	Raw *raw
}

var (
	// 数据库操作
	Crud *crud
)

// Config mysql配置
type Config struct {
	Host     string `env:"MYSQL_HOST"`
	Port     string `env:"MYSQL_PORT"`
	User     string `env:"MYSQL_USER"`
	Pass     string `env:"MYSQL_PASS"`
	Database string `env:"MYSQL_DATABASE"`

	// gorm.Config
	// 自定义logger
	Logger logger.Interface
	// 使用默认logger时，指定log level，默认为info
	LogLevel logger.LogLevel
	// 迁移模型时，不创建外键约束
	DisableForeignKeyConstraintWhenMigrating bool
	// table name prefix
	TablePrefix string

	// mysql.Config
	// 设置为0，string类型默认longtext
	// 否则，string类型默认varchar(DefaultStringSize)
	DefaultStringSize uint
}

// NewMysqlDB 创建mysql db
func NewMysqlDB(c Config) (*gorm.DB, error) {
	dialector := NewMysqlDialector(c)

	opt := &gorm.Config{
		Logger:                                   NewLogger(c),
		DisableForeignKeyConstraintWhenMigrating: c.DisableForeignKeyConstraintWhenMigrating,
		// tables, columns naming strategy
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: c.TablePrefix,
		},
	}
	db, err := gorm.Open(dialector, opt)
	return db, err
}

// NewLogger 创建logger
func NewLogger(c Config) logger.Interface {
	var l logger.Interface
	if c.Logger != nil {
		l = c.Logger
	} else {
		// 使用默认logger，log level为warn
		if c.LogLevel != 0 {
			// 外部设置log level
			l = logger.Default.LogMode(c.LogLevel)
		} else {
			// 默认log level为info
			l = logger.Default.LogMode(logger.Info)
		}
	}
	return l
}

// NewMysqlDialector 创建mysql dialector
func NewMysqlDialector(c Config) gorm.Dialector {
	// 当 SELECT 字段中有 DATE() 函数返回值时
	// parseTime=True 2023-10-10T00:00:00+08:00
	// parseTime=False 2023-10-10
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Pass, c.Host, c.Port, c.Database,
	)
	config := mysql.Config{
		DSN:               dsn,
		DefaultStringSize: c.DefaultStringSize,
	}
	dialector := mysql.New(config)
	return dialector
}

// AutoMigrate 自动迁移模型，不会删除列
// https://gorm.io/docs/migration.html
// opt是创建表选项 COMMENT='表注释'
// dst是数据模型指针
func (c *crud) AutoMigrate(opt string, dst ...any) error {
	if opt != "" {
		return c.db.Set("gorm:table_options", opt).AutoMigrate(dst...)
	} else {
		return c.db.AutoMigrate(dst...)
	}
}
