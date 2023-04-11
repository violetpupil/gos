package resty

import (
	"github.com/go-resty/resty/v2"
)

// Client http客户端
var Client *resty.Client

// Init 初始化客户端
func Init() {
	Client = resty.New()
}

// IsProxySet 客户端是否设置了代理
// 默认会使用环境变量的代理，参考 http.ProxyFromEnvironment
func IsProxySet() bool {
	if Client == nil {
		return false
	}
	return Client.IsProxySet()
}

// RemoveProxy 移除代理配置
func RemoveProxy() {
	if Client == nil {
		return
	}
	Client = Client.RemoveProxy()
}

// SetProxy 配置客户端代理 http://user:password@proxyserver:8888
// 默认会使用环境变量的代理，参考 http.ProxyFromEnvironment
// 会自动初始化客户端
func SetProxy(proxyURL string) {
	if Client == nil {
		Init()
	}
	Client = Client.SetProxy(proxyURL)
}
