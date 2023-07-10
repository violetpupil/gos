package zero

import (
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	// RpcServerConf is a rpc server config.
	// ListenOn 监听地址
	// Etcd etcd related configuration
	// Timeout 接口最大响应时间
	RpcServerConf = zrpc.RpcServerConf
	// RpcClientConf is a rpc client config.
	// Etcd etcd related configuration
	// Target rpc服务地址 localhost:9090
	RpcClientConf = zrpc.RpcClientConf
)

// Dial creates a client connection to the given target.
// 目标 localhost:9090
// 连接用完调用grpc.ClientConn.Close()关闭
func Dial(target string) (*grpc.ClientConn, error) {
	conf := zrpc.RpcClientConf{Target: target}
	zc, err := zrpc.NewClient(conf)
	if err != nil {
		logrus.Errorln("new client error", err)
		return nil, err
	}

	conn := zc.Conn()
	return conn, nil
}
