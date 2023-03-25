// https://github.com/nhooyr/websocket/tree/master/examples
package websocket

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/std/bufio"
	"nhooyr.io/websocket"
)

// Chat 和websocket服务通信
// 输入q退出聊天
// TODO 改为用ctrl-c退出，现在ctrl-c后的执行有不确定性，不能保证正常关闭连接
// TODO 接收消息
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

	logrus.Info("welcome join chat!")
	// 循环发送消息
	bufio.Scan(c.Stdin)

	err = conn.Close(websocket.StatusNormalClosure, "")
	logrus.WithField("Error", err).Info("close conn")
}

// Stdin 终端输入处理，返回是否退出bufio.Scan循环
func (c *Chat) Stdin(s string, err error) bool {
	if err != nil {
		// 循环已经结束
		logrus.Error("scan error ", err)
		return true
	}

	if s == "" {
		logrus.Info("cannot send empty")
		return false
	}
	// 用户退出聊天
	if s == "q" {
		logrus.Info("bye!")
		return true
	}
	err = c.conn.Write(c.ctx, websocket.MessageText, []byte(s))
	if err != nil {
		logrus.Error("write error ", err)
		return true
	}
	return false
}
