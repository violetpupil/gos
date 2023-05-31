// https://go-zero.dev/cn/
// https://go-zero.dev/cn/docs/introduction
// https://go-zero.dev/cn/docs/blog/blog
// https://github.com/zeromicro/zero-examples
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
		Host string
		Port int
		Pass string
	}
}
