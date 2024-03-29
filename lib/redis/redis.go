// https://redis.uptrace.dev/guide/
// 封装 redis 客户端，基于命令封装第三方库
// 必须先调用 Init 进行初始化
package redis

import (
	"github.com/redis/go-redis/v9"
)

// Nil 当键不存在时，有的命令会返回 Nil 错误
var Nil = redis.Nil

var client *redis.Client

// Config redis连接配置
type Config struct {
	Addr     string // 传空的话，默认连接localhost:6379
	Password string // 密码
	DB       int    // 数据库号
}

// Init 初始化客户端
// 初始化时不会连接 redis
func Init(c Config) {
	client = NewClient(c)
}

// NewClient 创建redis客户端
func NewClient(cfg Config) *redis.Client {
	opt := &redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	}
	c := redis.NewClient(opt)
	return c
}
