package redis

import (
	"context"
	"time"
)

// Incr 将指定键的值加1，返回增加后的值
// 键不存在的话，先创建键并设置为0，后执行操作
// https://redis.io/commands/incr/
func Incr(ctx context.Context, key string) (int64, error) {
	return client.Incr(ctx, key).Result()
}

// Decr 将指定键的值减1，返回减小后的值
// 键不存在的话，先创建键并设置为0，后执行操作
// https://redis.io/commands/decr/
func Decr(ctx context.Context, key string) (int64, error) {
	return client.Decr(ctx, key).Result()
}

// Get 获取字符串键值
// 键不存在的话，返回 redis.Nil 错误
// https://redis.io/commands/get/
func Get(ctx context.Context, key string) (string, error) {
	return client.Get(ctx, key).Result()
}

// Set 设置字符串，0代表不过期
// https://redis.io/commands/set/
func Set(ctx context.Context, key string, value interface{}, e time.Duration) error {
	return client.Set(ctx, key, value, e).Err()
}

// SetNx 设置字符串 如果不存在 0代表不过期
// 字符串存在返回redis.Nil
func SetNx(ctx context.Context, key string, value interface{}, e time.Duration) error {
	return client.SetNX(ctx, key, value, e).Err()
}

// GetDel 获取字符串键值，然后删除
// 键不存在的话，返回 redis.Nil 错误
// https://redis.io/commands/getdel/
func GetDel(ctx context.Context, key string) (string, error) {
	return client.GetDel(ctx, key).Result()
}
