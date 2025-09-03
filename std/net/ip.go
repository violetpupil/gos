package net

import (
	"errors"
	"net"

	"github.com/sirupsen/logrus"
)

// HostIp 获取主机ip
// 8.8.8.8 和 8.8.4.4 是谷歌提供的dns服务器
func HostIp() (net.IP, error) {
	conn, err := net.Dial("ip:icmp", "8.8.8.8")
	if err != nil {
		logrus.Error("dial error ", err)
		return nil, err
	}
	defer conn.Close()

	addr, ok := conn.LocalAddr().(*net.IPAddr)
	if !ok {
		return nil, errors.New("not ip addr")
	}
	return addr.IP, nil
}
