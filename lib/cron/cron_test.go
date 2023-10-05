package cron

import (
	"fmt"
	"testing"
)

func TestAddFunc(t *testing.T) {
	Init()
	AddFunc(Every("10s"), func() {
		fmt.Println("hi")
	})
	select {}
}
