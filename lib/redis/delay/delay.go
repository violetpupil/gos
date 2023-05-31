// redis延迟队列
// https://github.com/ouqiang/delay-queue
package delay

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/spec/types"
)

const (
	// 使用list存放到期的任务id
	lName = "dq_queue"
	// 使用sorted set存放未到期的任务id
	zName = "dq_bucket"
)

var client *redis.Client

// Init 初始化redis客户端
// 启动goroutine，使用f处理到期的任务，参数为任务id数组
func Init(addr, pass string, f func([]string)) {
	opt := &redis.Options{
		Addr:     addr,
		Password: pass,
	}
	client = redis.NewClient(opt)
	// 每秒查询一次是否到期
	time.AfterFunc(time.Second, move)
	go bLPop(f)
}

// ZAdd 设置任务延迟，id为任务标识
func ZAdd(d time.Duration, id string) error {
	ctx := context.Background()
	ts := time.Now().Add(d).Unix()
	z := redis.Z{Score: float64(ts), Member: id}
	err := client.ZAdd(ctx, zName, z).Err()
	return err
}

// ZRem 移除任务检测
func ZRem(ids []string) error {
	ctx := context.Background()
	err := client.ZRem(ctx, zName, ids).Err()
	return err
}

// zRange 到期任务检测
func zRange() ([]string, error) {
	ctx := context.Background()
	ts := time.Now().Unix()
	by := &redis.ZRangeBy{Min: "-inf", Max: fmt.Sprintf("%d", ts)}
	ids, err := client.ZRangeByScore(ctx, zName, by).Result()
	return ids, err
}

// rPush 将任务id放到到期队列
func rPush(ids []string) error {
	ctx := context.Background()
	idsI := types.ToSlice(ids)
	err := client.RPush(ctx, lName, idsI...).Err()
	return err
}

// bLPop 处理到期的任务，f参数为任务id数组
// 该函数会阻塞
func bLPop(f func([]string)) {
	for {
		ctx := context.Background()
		ids, err := client.BLPop(ctx, 0, lName).Result()
		if err != nil {
			logrus.Errorln("block left pop error", err)
			time.Sleep(10 * time.Second)
			continue
		}
		f(ids)
	}
}

// move 将到期的任务id转移到list
func move() {
	ids, err := zRange()
	if err != nil {
		logrus.Errorln("zrange expire error", err)
		return
	}
	if len(ids) == 0 {
		logrus.Infoln("expire empty")
		return
	}

	err = rPush(ids)
	if err != nil {
		logrus.Errorln("push expire error", err)
		return
	}
	err = ZRem(ids)
	if err != nil {
		logrus.Errorln("rem expire error", err)
	}
}
