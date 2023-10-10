// https://gorm.io/
// https://gorm.io/docs/
package gorm

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type (
	// 数据库对象
	// Select() 指定字段
	// 查询
	// Select("name, age")
	// Select("name", "age")
	// Select([]string{"name", "age"})
	// 如果两个字段名相同，使用后面的一个
	DB = gorm.DB
)

// crud 数据库操作
type crud struct {
	db  *gorm.DB
	C   *create
	R   *query
	U   *update
	D   *delete
	T   *transaction
	Raw *raw
}

var (
	// 必须先调用 Init 初始化
	// 数据库对象
	D *gorm.DB
	// 数据库操作
	Crud *crud
)

// Init 初始化数据库实例
func Init(db *gorm.DB) {
	D = db
	Crud = new(crud)
	Crud.db = db
	Crud.C = &create{db}
	Crud.R = &query{db}
	Crud.U = &update{db}
	Crud.D = &delete{db}
	Crud.T = &transaction{db}
	Crud.Raw = &raw{db}
}

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

// InitMySQL 初始化mysql数据库操作
// https://gorm.io/docs/connecting_to_the_database.html
// https://github.com/go-sql-driver/mysql
func InitMySQL(c Config) error {
	db, err := NewMysqlDB(c)
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

// NewMysqlDBLife 创建mysql db
// sets the maximum amount of time a connection may be reused.
// If t <= 0, connections are not closed due to a connection's age.
func NewMysqlDBLife(c Config, t time.Duration) (*gorm.DB, error) {
	db, err := NewMysqlDB(c)
	if err != nil {
		logrus.Errorln("new mysql db error", err)
		return nil, err
	}

	d, err := db.DB()
	if err != nil {
		logrus.Errorln("get sql db error", err)
		return nil, err
	}
	d.SetConnMaxLifetime(t)
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
