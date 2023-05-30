package redis

import (
	"context"
	"time"
)

// BLPop block left pop
// A timeout of zero can be used to block indefinitely.
// 超时后返回redis.Nil错误
// https://redis.io/commands/blpop/
func BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return client.BLPop(ctx, timeout, keys...).Result()
}
