package client

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/gos/lib/grpc"
	"github.com/violetpupil/gos/lib/grpc/examples/user"
)

func GetUser(id string) (*user.Response, error) {
	// 创建连接
	conn, err := grpc.Dial("localhost:9090")
	if err != nil {
		logrus.Errorln("grpc dial error", err)
		return nil, err
	}
	defer conn.Close()

	// 超时取消
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 使用生成的客户端
	c := user.NewUserClient(conn)
	res, err := c.GetUser(ctx, &user.IdRequest{Id: id})
	return res, err
}
