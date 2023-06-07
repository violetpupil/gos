package redis

import "context"

// SAdd set add
// 返回添加成功的元素数
// https://redis.io/commands/sadd/
func SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return client.SAdd(ctx, key, members...).Result()
}

// SRem set rem
// 返回移除成功的元素数
// https://redis.io/commands/srem/
func SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return client.SRem(ctx, key, members...).Result()
}
