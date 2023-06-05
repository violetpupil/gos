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

// Delay redis延迟队列
type Delay struct {
	client *redis.Client
	// redis list 存放到期的任务id
	lName string
	// redis sorted set 存放未到期的任务id
	zName string
}

// NewDelay 创建Delay实例
// 启动goroutine，使用f处理到期的任务
// f参数第一个元素是list名，第二个元素是值
//
// name用于组成redis键名
// list 存放到期的任务id，键名为dq_queue:name
// sorted set 存放未到期的任务id，键名为dq_bucket:name
// 每次创建Delay实例时，有序集合键名应该在redis库空间中唯一
// 否则，可能导致list中的任务id重复
func NewDelay(r *redis.Client, name string, f func([]string)) *Delay {
	d := &Delay{
		client: r,
		lName:  "dq_queue:" + name,
		zName:  "dq_bucket:" + name,
	}
	// 每秒查询一次是否到期
	time.AfterFunc(time.Second, d.move)
	go d.bLPop(f)
	return d
}

// ZAdd 设置任务延迟，id为任务标识
func (d *Delay) ZAdd(dur time.Duration, id string) error {
	ctx := context.Background()
	ts := time.Now().Add(dur).Unix()
	z := redis.Z{Score: float64(ts), Member: id}
	err := d.client.ZAdd(ctx, d.zName, z).Err()
	return err
}

// ZRem 移除未到期的任务
func (d *Delay) ZRem(ids []string) error {
	ctx := context.Background()
	err := d.client.ZRem(ctx, d.zName, ids).Err()
	return err
}

// RPush 将任务id放到到期队列
func (d *Delay) RPush(ids []string) error {
	ctx := context.Background()
	idsI := types.ToSlice(ids)
	err := d.client.RPush(ctx, d.lName, idsI...).Err()
	return err
}

// LMove 指定源队列及位置left/right，放到到期队列头
func (d *Delay) LMove(src, pos string) error {
	ctx := context.Background()
	err := d.client.LMove(ctx, src, d.lName, pos, "left").Err()
	return err
}

// zRange 到期任务检测
func (d *Delay) zRange() ([]string, error) {
	ctx := context.Background()
	ts := time.Now().Unix()
	by := &redis.ZRangeBy{Min: "-inf", Max: fmt.Sprintf("%d", ts)}
	ids, err := d.client.ZRangeByScore(ctx, d.zName, by).Result()
	return ids, err
}

// bLPop 处理到期的任务
// f参数第一个元素是list名，第二个元素是值
// 该函数会阻塞
func (d *Delay) bLPop(f func([]string)) {
	for {
		ctx := context.Background()
		id, err := d.client.BLPop(ctx, 0, d.lName).Result()
		if err != nil {
			logrus.Errorln("block pop error", err)
			time.Sleep(10 * time.Second)
			continue
		}
		f(id)
	}
}

// move 将到期的任务id转移到list
func (d *Delay) move() {
	ids, err := d.zRange()
	if err != nil {
		logrus.Errorln("zrange expire error", err)
		return
	}
	if len(ids) == 0 {
		logrus.Infoln("expire empty")
		return
	}

	err = d.RPush(ids)
	if err != nil {
		logrus.Errorln("push expire error", err)
		return
	}
	err = d.ZRem(ids)
	if err != nil {
		logrus.Errorln("rem expire error", err)
	}
}
