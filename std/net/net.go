package net

import (
	"errors"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)

// Conn 网络连接
type Conn = net.Conn

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

// Private 如果是私有ip，转为4字节表示返回，否则返回nil
// 私有ip地址范围
// 10.0.0.0 ~ 10.255.255.255
// 172.16.0.0 ~ 172.31.255.255
// 192.168.0.0 ~ 192.168.255.255
func Private(ip net.IP) net.IP {
	ip = ip.To4()
	if ip == nil {
		return nil
	}

	if ip[0] == 10 ||
		ip[0] == 172 && ip[1] >= 16 && ip[1] < 32 ||
		ip[0] == 192 && ip[1] == 168 {
		return ip
	}
	return nil
}

// ConnId 连接id，网络名-本地主机端口-远程主机端口
func ConnId(c net.Conn) string {
	l := c.LocalAddr()
	r := c.RemoteAddr()
	if l.Network() == r.Network() {
		return fmt.Sprintf("%s-%s-%s", l.Network(), l, r)
	} else {
		return fmt.Sprintf("%s-%s-%s-%s", l.Network(), l, r.Network(), r)
	}
}
