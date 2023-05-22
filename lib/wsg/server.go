package wsg

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// Upgrader http升级到websocket
var Upgrader = websocket.Upgrader{}

// Upgrade 升级协议，实现http.HandlerFunc函数类型
func Upgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorln("upgrade websocket error", err)
		return
	}
	defer conn.Close()

	for {
		t, p, e := conn.ReadMessage()
		if websocket.IsCloseError(e, websocket.CloseNormalClosure) {
			logrus.Infoln("client normal close")
			break
		}
		if e != nil {
			logrus.Errorln("read message error", e)
			break
		}

		logrus.Infof("read message %s", p)
		e = conn.WriteMessage(t, p)
		if e != nil {
			logrus.Errorln("write message error", e)
			break
		}
	}
}
