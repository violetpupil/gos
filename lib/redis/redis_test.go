package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	Init(Config{
		Addr: "localhost:6379",
	})
	status := client.Ping(context.Background())
	fmt.Println(status.String())
}
