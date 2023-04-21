// https://en.wikipedia.org/wiki/Cron
package cron

import (
	"context"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

// 调度器，必须先调用Start启动
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
	logrus.Infoln("add func entry %+v", entry)
	return nil
}

// AddFuncs 添加多个调度函数
func AddFuncs(specs map[string]func()) error {
	for k, v := range specs {
		err := AddFunc(k, v)
		if err != nil {
			logrus.Errorln("add funcs error", err)
			return err
		}
	}
	return nil
}
