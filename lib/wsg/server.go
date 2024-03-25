package wsg

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

// Upgrader http升级到websocket
var Upgrader = websocket.Upgrader{}

func GinUpgrade(c *gin.Context) {
	conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// If the upgrade fails, then Upgrade replies to the client with an HTTP error response.
		zap.L().Error("upgrade error", zap.Error(err))
		return
	}
	defer conn.Close()

	zap.L().Info("upgrade websocket conn",
		zap.String("localAddr", conn.LocalAddr().String()),
		zap.String("remoteAddr", conn.RemoteAddr().String()),
	)

	// 处理连接
}

// Upgrade 升级协议，实现http.HandlerFunc函数类型
func Upgrade(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("remoteAddr", r.RemoteAddr)
	logger.Infoln("upgrade websocket request")
	// If the upgrade fails, then Upgrade replies to the client with an HTTP error response.
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Errorln("upgrade websocket error", err)
		return
	}
	defer conn.Close()
	logrus.WithFields(logrus.Fields{
		"localAddr":  conn.LocalAddr(),
		"remoteAddr": conn.RemoteAddr(),
	}).Infoln("upgrade websocket conn")

	Echo(conn)
}

// Echo 回显消息
func Echo(conn *websocket.Conn) {
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
