// redis延迟队列
// https://github.com/ouqiang/delay-queue
package delay

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	// 使用list存放到期的任务id
	LName = "dq_queue"
	// 使用sorted set存放未到期的任务id
	ZName = "dq_bucket"
)

var client *redis.Client

// Init 初始化客户端
// 初始化时不会连接 redis
func Init(addr, pass string) {
	opt := &redis.Options{
		Addr:     addr,
		Password: pass,
	}
	client = redis.NewClient(opt)
	// 每秒查询一次是否到期
	time.AfterFunc(time.Second, Move)
}

// ZAdd 设置任务延迟，id为任务标识
func ZAdd(d time.Duration, id string) error {
	ctx := context.Background()
	ts := time.Now().Add(d).Unix()
	z := &redis.Z{Score: float64(ts), Member: id}
	err := client.ZAdd(ctx, ZName, z).Err()
	return err
}

// ZRange 获取到期的任务
func ZRange() ([]string, error) {
	ctx := context.Background()
	ts := time.Now().Unix()
	by := &redis.ZRangeBy{Min: "-inf", Max: fmt.Sprintf("%d", ts)}
	ids, err := client.ZRangeByScore(ctx, ZName, by).Result()
	return ids, err
}

// Move 将到期的任务id转移到list
func Move() {}
