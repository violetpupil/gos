package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestBLPop(t *testing.T) {
	Init(Config{})
	ctx := context.Background()
	r, err := BLPop(ctx, 0, "tmp")
	fmt.Println(r, err)
}

func TestRPush(t *testing.T) {
	Init(Config{})
	ctx := context.Background()
	r, err := RPush(ctx, "tmp", 1)
	fmt.Println(r, err)
}
