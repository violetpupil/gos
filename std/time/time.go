package time

import (
	"log"
	"time"
)

// 最新go版本中已支持这三个常量
const (
	DateTime = "2006-01-02 15:04:05"
	DateOnly = "2006-01-02"
	TimeOnly = "15:04:05"
)

// Cost 标准日志库记录函数耗时，在函数顶部调用 defer Cost()()
func Cost() func() {
	s := time.Now()
	f := func() {
		log.Print("cost: ", time.Since(s))
	}
	return f
}

// In 修改时间时区
func In(t time.Time, utcOffset float32) time.Time {
	return t.In(FixedZone(utcOffset))
}

// FixedZone 返回指定时区的 Location 对象
func FixedZone(utcOffset float32) *time.Location {
	return time.FixedZone("", int(utcOffset*3600))
}
