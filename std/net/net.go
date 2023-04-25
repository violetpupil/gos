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

// Listen 监听请求，并在单独的goroutine中处理连接
// The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
// For TCP networks, if the host in the address parameter is empty or a literal unspecified IP address,
// Listen listens on all available unicast and anycast IP addresses of the local system.
// tcp -> [::]
// tcp4 -> 0.0.0.0
//
// If the port in the address parameter is empty or "0",
// as in "127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
//
// address 取值 host:port 空字符串 host: :port
func Listen(network, address string, f func(net.Conn)) error {
	ln, err := net.Listen(network, address)
	if err != nil {
		logrus.Errorln("listen error", err)
		return err
	}
	logrus.Infoln("listener on", ln.Addr())

	for {
		// 阻塞等待连接
		conn, err := ln.Accept()
		if err != nil {
			logrus.WithField("ConnId", ConnId(conn)).Errorln("accept error", err)
			continue
		}
		logrus.WithField("ConnId", ConnId(conn)).Infoln("accept conn")
		go f(conn)
	}
}
