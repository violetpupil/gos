// Redis-based distributed mutual exclusion lock
// https://redis.io/docs/manual/patterns/distributed-locks/
package redsync

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

// New 创建redis分布式锁管理
func New(client *redis.Client) *redsync.Redsync {
	pool := goredis.NewPool(client)
	m := redsync.New(pool)
	return m
}
