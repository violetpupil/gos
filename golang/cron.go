package golang

import (
	"time"

	"github.com/robfig/cron/v3"
)

func NewCron() *cron.Cron {
	return cron.New(cron.WithLocation(time.FixedZone("Asia/Shanghai", 8*3600)))
}
