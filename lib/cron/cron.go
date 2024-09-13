package cron

import (
	"context"
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

const (
	// cron表达式
	Minutely = "* * * * *" // 每分钟
	Hourly   = "@hourly"   // 每小时
	Weekly   = "@weekly"   // 每周
	Monthly  = "@monthly"  // 每月
)

// Every 指定间隔时间 1ms -1s 1h0.5m
// 参考time.ParseDuration函数
func Every(duration string) string {
	return fmt.Sprintf("@every %s", duration)
}

// Daily 每天几点
// @daily 不准，可能一天执行两次
// 有时候在00:00:00执行
// 有时候在23:59:59执行
func Daily(hour int) string {
	return fmt.Sprintf("0 %d * * *", hour)
}

// 调度器
// Start() 在单独的goroutine启动，已经启动则无动作
// Run() 阻塞启动
var c *cron.Cron

// Init 初始化调度器
func Init() {
	c = cron.New()
	c.Start()
}

// Stop stops the cron scheduler if it is running; otherwise it does nothing.
// does not stop any jobs already running
// A context is returned so the caller can wait for running jobs to complete.
func Stop() context.Context {
	return c.Stop()
}

// AddFunc adds a func to the Cron to be run on the given schedule.
// Cron will run them in their own goroutines.
// 不会立即运行
func AddFunc(spec string, cmd func()) error {
	id, err := c.AddFunc(spec, cmd)
	if err != nil {
		logrus.Errorln("add func error", err)
		return err
	}
	entry := c.Entry(id)
	logrus.Infof("add func entry %+v", entry)
	return nil
}

// AddFuncNow adds a func to the Cron to be run on the given schedule.
// Cron will run them in their own goroutines.
// 会立即运行一次
func AddFuncNow(spec string, cmd func()) error {
	id, err := c.AddFunc(spec, cmd)
	if err != nil {
		logrus.Errorln("add func error", err)
		return err
	}
	entry := c.Entry(id)
	logrus.Infof("add func entry %+v", entry)

	go cmd()
	return nil
}
