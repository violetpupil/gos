package imap

import (
	"github.com/emersion/go-imap/client"
	"github.com/violetpupil/gos/lib/logrus"
)

func Login(addr, username, password string) error {
	conn, err := client.DialTLS(addr, nil)
	if err != nil {
		logrus.Errorln("dial tls error", err)
		return err
	}
	err = conn.Login(username, password)
	return err
}
