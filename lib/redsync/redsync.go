// Redis-based distributed mutual exclusion lock
// https://redis.io/docs/manual/patterns/distributed-locks/
package redsync

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

type (
	Redsync = redsync.Redsync
)

// m redis分布式锁管理
// 必须先调用Init初始化
var m *redsync.Redsync

// Init 初始化锁管理
func Init(client *redis.Client) {
	pool := goredis.NewPool(client)
	m = redsync.New(pool)
}

// New 创建redis分布式锁管理
func New(client *redis.Client) *redsync.Redsync {
	pool := goredis.NewPool(client)
	m := redsync.New(pool)
	return m
}

// NewMutex returns a new distributed mutex with given name.
// 使用redsync.Mutex.Lock()和redsync.Mutex.Unlock()加解锁
func NewMutex(name string) *redsync.Mutex {
	return m.NewMutex(name)
}
