package imap

import (
	"github.com/emersion/go-imap/client"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/proxy"
)

// Login 登录邮箱
func Login(addr, username, password string) error {
	c, err := client.DialTLS(addr, nil)
	if err != nil {
		logrus.Errorln("dial tls error", err)
		return err
	}

	err = c.Login(username, password)
	if err != nil {
		logrus.Errorln("login error", err)
		return err
	}
	defer c.Logout()
	return nil
}

// LoginSocks5 用socks5代理登录邮箱
func LoginSocks5(addr, username, password, addrP, usernameP, passwordP string) error {
	c, err := DialTLSSocks5(addr, addrP, usernameP, passwordP)
	if err != nil {
		logrus.Errorln("dial tls socks5 error", err)
		return err
	}

	err = c.Login(username, password)
	if err != nil {
		logrus.Errorln("login error", err)
		return err
	}
	defer c.Logout()
	return nil
}

// DialTLSSocks5 连接imap服务器 使用socks5代理
func DialTLSSocks5(addr, addrP, usernameP, passwordP string) (*client.Client, error) {
	dialer, err := proxy.SOCKS5("tcp", addrP, &proxy.Auth{
		User:     usernameP,
		Password: passwordP,
	}, proxy.Direct)
	if err != nil {
		logrus.Errorln("socks5 dialer error", err)
		return nil, err
	}

	c, err := client.DialWithDialerTLS(dialer, addr, nil)
	return c, err
}
