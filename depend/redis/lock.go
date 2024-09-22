package redis

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

func NewLockManager(c *redis.Client) *redsync.Redsync {
	pool := goredis.NewPool(c)
	return redsync.New(pool)
}
