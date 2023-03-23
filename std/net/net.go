package net

import (
	"errors"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/std/os"
)

// 接口
type (
	// Conn 网络连接
	Conn = net.Conn
)

// 结构体
type (
	// OpError 读写网络等操作时异常
	OpError = net.OpError
)

// LogOpError 打印net.OpError相关信息
// 如果不是net.OpError，直接返回
func LogOpError(err error) {
	opErr := new(net.OpError)
	if !errors.As(err, &opErr) {
		return
	}

	errWrap := errors.Unwrap(opErr)
	logrus.WithFields(logrus.Fields{
		"Error":         opErr,
		"Temporary":     opErr.Temporary(),
		"Timeout":       opErr.Timeout(),
		"ErrorWrapType": fmt.Sprintf("%T", errWrap),
	}).Errorf("%+v", *opErr)

	os.LogSyscallError(errWrap)
}

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
	if c == nil {
		return ""
	}

	l := c.LocalAddr()
	r := c.RemoteAddr()
	if l.Network() == r.Network() {
		return fmt.Sprintf("%s-%s-%s", l.Network(), l, r)
	} else {
		return fmt.Sprintf("%s-%s-%s-%s", l.Network(), l, r.Network(), r)
	}
}
