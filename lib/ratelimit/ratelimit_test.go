package ratelimit

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/yudeguang/ratelimit"
)

func TestNewRule(t *testing.T) {
	r := ratelimit.NewRule()
	r.AddRule(time.Second, 10)
	for i := 0; i <= 20; i++ {
		allow := r.AllowVisit("user")
		if !allow {
			log.Println("访问量超出,其剩余访问次数情况如下:", r.RemainingVisits("user"))
		} else {
			log.Println("允许访问,其剩余访问次数情况如下:", r.RemainingVisits("user"))
		}
	}
	time.Sleep(time.Second)
	fmt.Println(r.AllowVisit("user"))
}
