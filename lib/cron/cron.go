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
// @monthly @weekly @daily @hourly
// @every <duration> 指定间隔时间执行
//
// 用,指定多个时间，用*/指定间隔时间执行
// */5 1,2,3 * * * echo hello world
package cron

import (
	"context"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

const (
	// cron表达式
	// 每分钟
	Minutely = "* * * * *"
)

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
