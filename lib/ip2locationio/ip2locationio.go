// https://ip2location-io-go.readthedocs.io/en/latest/
package ip2locationio

import (
	"github.com/ip2location/ip2location-io-go/ip2locationio"
	"github.com/sirupsen/logrus"
)

// 必须调用Init初始化
var IPGeolocation *ip2locationio.IPGeolocation

func Init(apiKey string) error {
	config, err := ip2locationio.OpenConfiguration(apiKey)
	if err != nil {
		logrus.Errorln("open configuration error", err)
		return err
	}
	IPGeolocation, err = ip2locationio.OpenIPGeolocation(config)
	return err
}

// LookUp 查询ip位置信息
func LookUp(ip, lang string) (ip2locationio.IPGeolocationResult, error) {
	return IPGeolocation.LookUp(ip, lang)
}
