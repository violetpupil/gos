// ┌───────────── minute (0 - 59)
// │ ┌───────────── hour (0 - 23)
// │ │ ┌───────────── day of the month (1 - 31)
// │ │ │ ┌───────────── month (1 - 12)
// │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday;
// │ │ │ │ │                                   7 is also Sunday on some systems)
// │ │ │ │ │
// │ │ │ │ │
// * * * * * <command to execute>
//
// 用,指定多个时间，用*/指定间隔时间执行
// */5 1,2,3 * * * echo hello world
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
	Daily    = "@daily"    // 每天
	Weekly   = "@weekly"   // 每周
	Monthly  = "@monthly"  // 每月
)

// Every 指定间隔时间 1ms -1s 1h0.5m
// 参考time.ParseDuration函数
func Every(duration string) string {
	return fmt.Sprintf("@every %s", duration)
}

// 调度器，必须先调用Start启动
// Start() 在goroutine启动
// Run() 阻塞启动
var c *cron.Cron

// Start the cron scheduler in its own goroutine, or no-op if already started.
func Start() {
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

	cmd()
	return nil
}
