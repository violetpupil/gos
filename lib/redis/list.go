package redis

import (
	"context"
	"time"
)

// BLPop 阻塞弹出第一个元素
// 如果多个键都有元素，只弹出指定顺序靠前的一个键
// 返回切片，第一个元素是list名，第二个元素是值
// A timeout of zero can be used to block indefinitely.
// 超时后返回redis.Nil错误
// https://redis.io/commands/blpop/
func BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return client.BLPop(ctx, timeout, keys...).Result()
}

// LPop 弹出第一个元素
// key不存在时，返回redis.Nil错误
// https://redis.io/commands/lpop/
func LPop(ctx context.Context, key string) (string, error) {
	return client.LPop(ctx, key).Result()
}

// RPush 在尾部添加元素
// Return the length of the list after the push operation.
// https://redis.io/commands/rpush/
func RPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return client.RPush(ctx, key, values...).Result()
}

// LMove 移动元素
// pos为left或right
// https://redis.io/commands/lmove/
func LMove(ctx context.Context, source, destination, srcPos, destPos string) (string, error) {
	return client.LMove(ctx, source, destination, srcPos, destPos).Result()
}

// LRem 移除元素
// Return the number of removed elements.
// count > 0: Remove elements equal to element moving from head to tail.
// count < 0: Remove elements equal to element moving from tail to head.
// count = 0: Remove all elements equal to element.
// https://redis.io/commands/lrem/
func LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error) {
	return client.LRem(ctx, key, count, value).Result()
}
