// https://github.com/nhooyr/websocket/tree/master/examples
package websocket

import (
	"context"

	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
)

type Chat struct {
	ctx  context.Context
	conn *websocket.Conn
}

// Chat 和websocket服务通信
// URLs with http/https schemes will work and are interpreted as ws/wss.
func (c *Chat) Chat(url string) {
	c.ctx = context.Background()
	conn, _, err := websocket.Dial(c.ctx, url, nil)
	if err != nil {
		logrus.Fatal("dial error ", err)
	}
	c.conn = conn
	// 如果已经关闭，则无操作
	defer conn.Close(websocket.StatusInternalError, "the sky is falling")

	err = c.conn.Write(c.ctx, websocket.MessageText, []byte("hi"))
	if err != nil {
		logrus.Fatal("write error ", err)
	}
	t, bs, err := c.conn.Read(c.ctx)
	logrus.WithFields(logrus.Fields{
		"Type":    t,
		"Message": string(bs),
		"Error":   err,
	}).Info("read websocket")

	err = conn.Close(websocket.StatusNormalClosure, "")
	logrus.WithField("Error", err).Info("close conn")
}
