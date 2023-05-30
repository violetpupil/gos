package redis

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestExpire(t *testing.T) {
	Init(Config{})
	ctx := context.Background()
	r, e := Expire(ctx, "tmp", time.Second, "")
	fmt.Println(r, e)
}
