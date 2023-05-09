package redis

import "context"

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
// 键不存在的话，返回 Nil 错误
// https://redis.io/commands/get/
func Get(ctx context.Context, key string) (string, error) {
	return client.Get(ctx, key).Result()
}
