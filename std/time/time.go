// https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
package time

import (
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	// 视频时间
	VideoTime = "15:04:05,000"
)

const (
	Second = time.Second
)

// 函数
var (
	Now   = time.Now
	Since = time.Since
	Sleep = time.Sleep
	// AfterFunc waits for the duration to elapse and then calls f in its own goroutine.
	AfterFunc = time.AfterFunc
	// ParseDuration parses a duration string.
	// 1ms -1s 1h0.5m
	ParseDuration = time.ParseDuration
)

type (
	Duration = time.Duration
	// A Ticker holds a channel that delivers “ticks” of a clock at intervals.
	// The ticker will adjust the time interval or drop ticks to make up for slow receivers.
	// Stop the ticker to release associated resources.
	Ticker = time.Ticker
	// 单次计时器
	Timer = time.Timer
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
	return time.UnixMilli(msec).UTC().Format(VideoTime)
}

// DateTime 将秒级时间戳转为2006-01-02 15:04:05
func DateTime(ts int64) string {
	return time.Unix(ts, 0).Format(time.DateTime)
}

// NewLoop 启动goroutine，一直执行f，并间隔一段时间
func NewLoop(d time.Duration, f func()) {
	go func() {
		for {
			f()
			time.Sleep(d)
		}
	}()
}

// NowStr 当前时间字符串 20230909191209572
func NowStr() string {
	s := time.Now().Format("20060102150405.000")
	s = strings.ReplaceAll(s, ".", "")
	return s
}

// ElapseMonth 指定时间是否过了一个月 24*30小时
func ElapseMonth(t time.Time) bool {
	return time.Since(t) > 24*30*time.Hour
}
