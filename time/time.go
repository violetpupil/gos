package time

import (
	"log"
	"time"
)

// Beauty 好看的时间格式
const Beauty = "2006-01-02 15:04:05"

// Cost 标准日志库记录函数耗时，在函数顶部调用 defer Cost()()
func Cost() func() {
	s := time.Now()
	f := func() {
		log.Print("cost: ", time.Since(s))
	}
	return f
}
