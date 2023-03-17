// websocket
// https://github.com/gobwas/ws-examples
package ws

import (
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/std/net"
)

// Websocket http处理器，访问任何地址都升级到websocket
// 同时为连接开启goroutine，回显消息
type Websocket struct{}

// ServeHTTP 实现http.Handler
func (s *Websocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		logrus.Error("upgrade websocket error ", err)
		w.WriteHeader(http.StatusInternalServerError)
		cnt, err := w.Write([]byte("sorry"))
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"Count": cnt,
				"Error": err,
			}).Error("response write")
		}
	}
	logrus.Info("upgrade websocket success ", net.ConnId(conn))

	go Echo(conn)
}

// Echo 回声
func Echo(conn net.Conn) {
	defer func() {
		logrus.Info("conn close ", net.ConnId(conn))
		conn.Close()
	}()

	// TODO 处理客户端断开，退出循环
	for {
		// 阻塞读取数据
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			logrus.Error("read client data error ", err)
			continue
		}
		err = wsutil.WriteServerMessage(conn, op, msg)
		logrus.WithFields(logrus.Fields{
			"Message": msg,
			"Opera":   op,
			"Error":   err,
		}).Info("write server message")
	}
}
