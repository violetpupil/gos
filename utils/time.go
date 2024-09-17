package utils

import "time"

func SetTimeZone(t time.Time, offset float32) time.Time {
	return t.In(time.FixedZone("", int(offset*3600)))
}
