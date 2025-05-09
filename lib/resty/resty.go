package resty

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// Client http客户端
// SetDebug() Client logs details of every request and response.
// SetDebugBodyLimit() log请求/响应体最大长度
//
// RemoveProxy() 移除代理配置
// IsProxySet() 客户端是否设置了代理
// 默认会使用环境变量的代理，参考 http.ProxyFromEnvironment
//
// R() 创建请求实例
// SetRetryCount() 启用重试并设置次数
// SetTimeout() 设置超时时间
var Client *resty.Client

// Init 初始化客户端
func Init() {
	Client = resty.New()
}

// SetProxy 配置客户端代理 http://user:password@proxyserver:8888
// 默认会使用环境变量的代理，参考 http.ProxyFromEnvironment
// 会自动初始化客户端 支持socks5
func SetProxy(proxyURL string) {
	if Client == nil {
		Init()
	}
	Client = Client.SetProxy(proxyURL)
}

// SetProxySct 配置客户端代理，使用url.URL结构体
// 默认会使用环境变量的代理，参考 http.ProxyFromEnvironment
// 会自动初始化客户端 支持socks5
func SetProxySct(u *url.URL) {
	proxyURL := u.String()
	SetProxy(proxyURL)
}

// SetProxyPart 配置客户端代理 http://proxyserver:8888, user, password
// 默认会使用环境变量的代理，参考 http.ProxyFromEnvironment
// 会自动初始化客户端 支持socks5
func SetProxyPart(uri, user, pass string) error {
	u, err := url.Parse(uri)
	if err != nil {
		logrus.Errorln("parse url error", err)
		return err
	}
	u.User = url.UserPassword(user, pass)
	SetProxySct(u)
	return nil
}

// SetProxyFive 配置客户端代理 scheme, host, port, user, password
// 默认会使用环境变量的代理，参考 http.ProxyFromEnvironment
// 会自动初始化客户端 支持socks5
func SetProxyFive(scheme, host, port, user, pass string) {
	SetProxy(fmt.Sprintf("%s://%s:%s@%s:%s", scheme, user, pass, host, port))
}
