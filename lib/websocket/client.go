package websocket

import (
	"context"

	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// Chat 和websocket服务通信
// URLs with http/https schemes will work and are interpreted as ws/wss.
func Chat(url string) {
	ctx := context.Background()
	c, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		logrus.Fatal("dial error ", err)
	}
	// 如果已经关闭，则无操作
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	err = wsjson.Write(ctx, c, "hi")
	if err != nil {
		logrus.Fatal("write error ", err)
	}

	err = c.Close(websocket.StatusNormalClosure, "")
	logrus.WithField("Error", err).Info("close conn")
}
