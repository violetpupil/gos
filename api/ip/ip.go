// ip信息api
package ip

import (
	"time"

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

// ProxyInterval 检测动态代理ip切换间隔
func ProxyInterval(uri, user, pass string) {
	var start time.Time
	var last string
	var times int

	for {
		time.Sleep(time.Second)
		this, err := ProxyIpEcho(uri, user, pass)
		if err != nil {
			logrus.Errorln("proxy ip echo error", err)
			continue
		}
		logrus.Infoln("proxy ip", this)

		if last == "" {
			last = this
			continue
		}
		// 从第一次改变开始计时
		if this != last {
			if times == 0 {
				start = time.Now()
				last = this
				times += 1
				continue
			}
			logrus.Infoln(time.Since(start))
			return
		}
	}
}
