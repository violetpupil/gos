// 封装 redis 客户端，基于命令封装第三方库
// 必须先调用 Init 进行初始化
package redis

import (
	"github.com/go-redis/redis/v8"
)

// Nil 当键不存在时，有的命令会返回 Nil 错误
var Nil = redis.Nil

var client *redis.Client

// Init 初始化客户端
// 初始化时不会连接 redis
// addr 传空的话，默认连接 localhost:6379
func Init(addr string) {
	opt := &redis.Options{
		Addr: addr,
	}
	client = redis.NewClient(opt)
}
