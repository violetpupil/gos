// The order of elements is from the lowest to the highest score.
// Elements with the same score are ordered lexicographically.
package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// ZRange 获取有序集合元素 根据索引
// start stop 前后包含
// zrange key 0 -1 获取全部元素
// Out of range indexes do not produce an error.
// https://redis.io/commands/zrange/
func ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return client.ZRange(ctx, key, start, stop).Result()
}

// ZRangeWithScores 获取有序集合元素 包含得分
func ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	return client.ZRangeWithScores(ctx, key, start, stop).Result()
}

// ZRangeByScore 获取有序集合元素 根据得分
// start stop 前后包含
// zrange key -inf +inf byScore 获取全部元素
func ZRangeByScore(ctx context.Context, key string, start, stop string) ([]string, error) {
	by := &redis.ZRangeBy{Min: start, Max: stop}
	return client.ZRangeByScore(ctx, key, by).Result()
}

// ZRem 移除元素值
// https://redis.io/commands/zrem/
func ZRem(ctx context.Context, key string, members ...interface{}) error {
	err := client.ZRem(ctx, key, members...).Err()
	return err
}
