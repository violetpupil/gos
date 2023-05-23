package wsg

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// Message websocket消息
type Message struct {
	Type int
	Data []byte
}

// Hub 连接中心
var Hub sync.Map

// WriteMessage 从Hub中获取websocket连接，写消息到channel
func WriteMessage(key string, t int, data []byte) error {
	i, ok := Hub.Load(key)
	if !ok {
		return errors.New("websocket connection not in hub")
	}
	c, ok := i.(*Conn)
	if !ok {
		return errors.New("hub value type mismatch")
	}

	msg := Message{t, data}
	c.Write <- msg
	return nil
}

// Conn websocket连接
type Conn struct {
	Conn *websocket.Conn
	// websocket.Conn support one concurrent reader and one concurrent writer.
	Write chan Message
}

// NewConn 保存连接到Hub
// 读取消息直到异常，f为消息处理函数，参数为消息类型和内容
// 创建goroutine，将Conn.Write中的消息写到连接
func NewConn(conn *websocket.Conn, key string, f func(int, []byte)) {
	// 读消息异常后，会使写消息goroutine退出
	// 先从Hub去掉，再关闭channel，防止panic
	c := &Conn{
		Conn:  conn,
		Write: make(chan Message),
	}
	defer close(c.Write)
	Hub.Store(key, c)
	defer Hub.Delete(key)

	// 写消息
	go func() {
		for msg := range c.Write {
			err := conn.WriteMessage(msg.Type, msg.Data)
			if err != nil {
				logrus.Errorln("write message error", err)
				continue
			}
			logrus.WithFields(logrus.Fields{
				"Type":    msg.Type,
				"Message": string(msg.Data),
			}).Infoln("write message")
		}
	}()

	// 读消息
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

		logrus.WithFields(logrus.Fields{
			"Type":    t,
			"Message": string(p),
		}).Infoln("read message")
		f(t, p)
	}
}
