// https://go-zero.dev/
// https://github.com/zeromicro/zero-examples
//
// go zero 目前支持k8s版本
// k8s.io/api v0.26.3
// k8s.io/apimachinery v0.27.0-alpha.3
// k8s.io/client-go v0.26.3
package zero

import "github.com/zeromicro/go-zero/zrpc"

// Config rpc服务配置
type Config struct {
	zrpc.RpcServerConf
	MySQL struct {
		Host     string
		Port     int
		User     string
		Pass     string
		Database string
	}
	// zrpc.RpcServerConf.Redis字段必须设置Key
	// 所以另外定义字段，且字段名不能和Redis重名
	RedisC struct {
		Addr string
		Pass string
		DB   int
	}
	Bool bool
}
