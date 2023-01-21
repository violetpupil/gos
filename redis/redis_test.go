package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestIncr(t *testing.T) {
	Init("")
	ctx := context.Background()
	r, e := Incr(ctx, "tmp")
	fmt.Print(r, e)
}
