package time

import (
	"fmt"
	"testing"
	"time"
)

func TestIn(t *testing.T) {
	fmt.Println(In(time.Now(), 0))
	fmt.Println(In(time.Now(), 8))
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		fmt.Println("hi")
	}
}

func TestNowStr(t *testing.T) {
	fmt.Println(NowStr())
}
