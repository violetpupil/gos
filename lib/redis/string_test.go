package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestIncr(t *testing.T) {
	Init(Config{})
	ctx := context.Background()
	r, e := Incr(ctx, "tmp")
	fmt.Println(r, e)
}

func TestDecr(t *testing.T) {
	Init(Config{})
	ctx := context.Background()
	r, e := Decr(ctx, "tmp")
	fmt.Println(r, e)
}

func TestGet(t *testing.T) {
	Init(Config{})
	ctx := context.Background()
	r, e := Get(ctx, "tmp")
	fmt.Println(r, e)
}
