// 封装 redis 客户端
// 必须先调用 Init 进行初始化
package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

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

// Incr 将指定键的值加1，返回增加后的值
// https://redis.io/commands/incr/
func Incr(ctx context.Context, key string) (int64, error) {
	return client.Incr(ctx, key).Result()
}
