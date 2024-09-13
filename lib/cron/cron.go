package cron

import (
	"context"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

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
