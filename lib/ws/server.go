// websocket
// https://www.rfc-editor.org/rfc/rfc6455
// https://github.com/gobwas/ws-examples
package ws

import (
	"errors"
	"fmt"
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
	// ws会构造响应
	conn, _, handshake, err := ws.UpgradeHTTP(r, w)
	log := logrus.WithFields(logrus.Fields{
		"ConnId":    net.ConnId(conn),
		"Handshake": fmt.Sprintf("%+v", handshake),
	})
	if err != nil {
		log.Error("upgrade websocket error ", err)
		return
	}
	log.Info("upgrade websocket success")

	go Echo(conn)
}

// Echo 回显消息
func Echo(conn net.Conn) {
	defer func() {
		err := conn.Close()
		logrus.WithFields(logrus.Fields{
			"ConnId": net.ConnId(conn),
			"Error":  err,
		}).Info("conn close")
	}()

	for {
		// 阻塞读取数据
		msg, op, err := wsutil.ReadClientData(conn)
		// 客户端发送关闭帧
		// 状态码1000代表正常关闭
		var close wsutil.ClosedError
		if errors.As(err, &close) {
			logrus.WithField("ConnId", net.ConnId(conn)).Info("client disconnect ", close)
			return
		}
		// 当发生异常时，有可能之后一直异常
		if err != nil {
			logrus.Error("read client data error ", err)
			return
		}

		err = wsutil.WriteServerMessage(conn, op, msg)
		if err != nil {
			logrus.Error("write server message error ", err)
			return
		}
		logrus.WithFields(logrus.Fields{
			"Message": string(msg),
			"Opera":   op,
		}).Info("write server message")
	}
}
