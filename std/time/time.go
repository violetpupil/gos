package time

import (
	"log"
	"strconv"
	"time"
)

const (
	// 最新go版本中已支持这三个常量
	DateTime = "2006-01-02 15:04:05"
	DateOnly = "2006-01-02"
	TimeOnly = "15:04:05"

	// 视频时间
	VideoTime = "15:04:05,000"
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

// Ts 当前秒级时间戳字符串
func Ts() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// Video 将毫秒数转为视频时间字符串
func Video(msec int64) string {
	return In(time.UnixMilli(msec), 0).Format(VideoTime)
}
