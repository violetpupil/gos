package imap

import (
	"github.com/emersion/go-imap/client"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/proxy"
)

// Login 登录邮箱
func Login(addr, username, password string) error {
	conn, err := client.DialTLS(addr, nil)
	if err != nil {
		logrus.Errorln("dial tls error", err)
		return err
	}
	err = conn.Login(username, password)
	return err
}

// LoginSocks5 用socks5代理登录邮箱
func LoginSocks5(addr, username, password, addrP, usernameP, passwordP string) error {
	dialer, err := proxy.SOCKS5("tcp", addrP, &proxy.Auth{
		User:     usernameP,
		Password: passwordP,
	}, proxy.Direct)
	if err != nil {
		logrus.Errorln("socks5 dialer error", err)
		return err
	}

	conn, err := client.DialWithDialerTLS(dialer, addr, nil)
	if err != nil {
		logrus.Errorln("dial tls error", err)
		return err
	}
	err = conn.Login(username, password)
	return err
}
