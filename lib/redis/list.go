package redis

import (
	"context"
	"time"
)

// BLPop block left pop
// A timeout of zero can be used to block indefinitely.
// 返回切片，第一个元素是list名，第二个元素是值
// 超时后返回redis.Nil错误
// https://redis.io/commands/blpop/
func BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return client.BLPop(ctx, timeout, keys...).Result()
}

// RPush right push
// Return the length of the list after the push operation.
// https://redis.io/commands/rpush/
func RPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return client.RPush(ctx, key, values...).Result()
}
