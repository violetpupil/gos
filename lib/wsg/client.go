package wsg

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// Chat 和websocket服务通信
func Chat(addr string) error {
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		logrus.Errorln("dial error", err)
		return err
	}
	defer conn.Close()

	err = conn.WriteMessage(websocket.TextMessage, []byte("hi"))
	if err != nil {
		logrus.Errorln("write message error", err)
		return err
	}
	_, message, err := conn.ReadMessage()
	if err != nil {
		logrus.Errorln("read message error", err)
		return err
	}
	logrus.Infof("read message %s", message)

	err = WriteClose(conn)
	return err
}
