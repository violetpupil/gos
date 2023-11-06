package main

import (
	"context"

	"github.com/violetpupil/gos/lib/grpc"
	"github.com/violetpupil/gos/lib/grpc/examples/user"
)

func main() {
	// 运行 grpc 服务
	grpc.Serve("9090", func(sr grpc.ServiceRegistrar) {
		// 注册用户服务
		user.RegisterUserServer(sr, &UserServer{})
	})
}

// UserServer 用户服务实现
type UserServer struct {
	user.UnimplementedUserServer
}

func (s *UserServer) GetUser(context.Context, *user.IdRequest) (*user.Response, error) {
	res := &user.Response{Code: user.ResponseCode_SUCCESS}
	return res, nil
}
