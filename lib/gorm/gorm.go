// https://gorm.io/
// https://gorm.io/docs/
package gorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
