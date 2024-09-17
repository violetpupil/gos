package utils

import "time"

func SetTimeZone(t time.Time, offset float32) time.Time {
	return t.In(time.FixedZone("", int(offset*3600)))
}

// 0时区今天剩余时间
func TodayLeft() time.Duration {
	now := time.Now().UTC()
	tomorrow := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, 1)
	return tomorrow.Sub(now)
}
