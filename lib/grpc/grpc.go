// https://grpc.io/docs/languages/go/
// https://github.com/grpc/grpc-go/tree/master/Documentation
// https://github.com/grpc/grpc-go/tree/master/examples
//
// The default logger is controlled by environment variables. Turn everything on like this:
// export GRPC_GO_LOG_VERBOSITY_LEVEL=99
// export GRPC_GO_LOG_SEVERITY_LEVEL=info
package grpc

import (
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// RegisterHealthServer 注册健康检查
func RegisterHealthServer(s grpc.ServiceRegistrar) {
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
}

// Serve 监听grpc请求，为每个连接创建一个goroutine
// f为注册函数
func Serve(port string, f func(grpc.ServiceRegistrar)) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logrus.Errorln("net listen error", err)
		return err
	}

	s := grpc.NewServer()
	if f != nil {
		f(s)
	}
	logrus.Infof("server listening at %v", lis.Addr())
	// Serve will return a non-nil error unless Stop or GracefulStop is called.
	err = s.Serve(lis)
	return err
}
