// ip信息api
package ip

import (
	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/lib/resty"
)

const (
	// 返回ip及地理信息
	LumTest = "https://lumtest.com/myip.json"
	// 只返回ip
	IpEcho = "http://ipecho.net/plain"
)

// 使用代理访问 ip echo
func ProxyIpEcho(uri, user, pass string) (string, error) {
	err := resty.SetProxyPart(uri, user, pass)
	if err != nil {
		logrus.Errorln("set proxy part error", err)
		return "", err
	}
	res, err := resty.Get(IpEcho, nil)
	if err != nil {
		logrus.Errorln("get error", err)
		return "", err
	}
	return res.String(), nil
}
