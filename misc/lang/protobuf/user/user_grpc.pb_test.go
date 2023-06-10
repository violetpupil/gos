package user

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/violetpupil/components/lib/grpc"
)

func Test_userClient_GetUser(t *testing.T) {
	conn, err := grpc.Dial("localhost:9090")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	c := NewUserClient(conn)
	res, err := c.GetUser(ctx, &IdRequest{Id: ""})
	fmt.Println(res, err)
}
