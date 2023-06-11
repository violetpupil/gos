package cron

import (
	"fmt"
	"testing"
)

func TestAddFunc(t *testing.T) {
	Start()
	AddFunc(Every("10s"), func() {
		fmt.Println("hi")
	})
	select {}
}
