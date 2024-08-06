// Redis-based distributed mutual exclusion lock
// https://redis.io/docs/manual/patterns/distributed-locks/
package redsync

import (
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// New 创建redis分布式锁管理
func New(client *redis.Client) *redsync.Redsync {
	pool := goredis.NewPool(client)
	m := redsync.New(pool)
	return m
}

// 加锁示例
func Lock() error {
	// prepare
	var sync *redsync.Redsync
	key := ""

	// 默认重试32次，每次重试间隔50~250ms
	lock := sync.NewMutex(
		key,
		redsync.WithExpiry(time.Minute),
		redsync.WithTries(1000),
	)
	err := lock.Lock()
	if err != nil {
		zap.L().Error("lock error", zap.Error(err))
		return err
	} else {
		zap.L().Info("lock success")
	}
	defer func() {
		r, err := lock.Unlock()
		zap.L().Info("unlock result", zap.Bool("status", r), zap.Error(err))
	}()

	// do something
	return nil
}
