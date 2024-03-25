package wsg

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// websocket消息类型
const (
	TextMessage   = websocket.TextMessage   // UTF-8 encoded text data
	BinaryMessage = websocket.BinaryMessage // binary data message
	CloseMessage  = websocket.CloseMessage  // close control message
	PingMessage   = websocket.PingMessage   // ping control message
	PongMessage   = websocket.PongMessage   // pong control message
)

// Message websocket消息
type Message struct {
	Type int
	Data []byte
}

// Hub 连接中心
var Hub sync.Map

// CheckConn 检查是否建立了websocket连接
func CheckConn(key string) bool {
	_, ok := Hub.Load(key)
	return ok
}

// WriteMessage 从Hub中获取websocket连接，写消息到channel
// 返回nil不代表向连接写消息成功
// t是websocket消息类型
func WriteMessage(key string, t int, data []byte) error {
	logger := zap.L().With(
		zap.String("key", key),
		zap.Int("type", t),
		zap.String("data", string(data)),
	)

	i, ok := Hub.Load(key)
	if !ok {
		logger.Info("websocket connection not in hub")
		return errors.New("websocket connection not in hub")
	}
	c, ok := i.(*Conn)
	if !ok {
		logger.Info("hub value type mismatch")
		return errors.New("hub value type mismatch")
	}

	msg := Message{t, data}
	c.Write <- msg
	return nil
}

// Conn websocket连接
type Conn struct {
	// The default close handler sends a close message to the peer.
	// The default ping handler sends a pong message to the peer.
	Conn *websocket.Conn
	// websocket.Conn support one concurrent reader and one concurrent writer.
	Write chan Message
}

func (c *Conn) Logger(key string) *zap.Logger {
	return zap.L().With(
		zap.String("key", key),
		zap.String("localAddr", c.Conn.LocalAddr().String()),
		zap.String("remoteAddr", c.Conn.RemoteAddr().String()),
	)
}

// NewConn 保存连接到Hub
// 读取消息直到异常，f为消息处理函数
// 创建goroutine，将Conn.Write中的消息写到连接
// 如果 Hub 中已经有 key 的连接，则报错
func NewConn(conn *websocket.Conn, key string, f func(key string, t int, data []byte)) error {
	// 读消息异常后，会使写消息goroutine退出
	// 先从Hub去掉，再关闭channel，防止panic
	c := &Conn{
		Conn:  conn,
		Write: make(chan Message),
	}
	defer close(c.Write)

	logger := c.Logger(key)

	_, ok := Hub.LoadOrStore(key, c)
	if ok {
		logger.Info("conn established")
		return errors.New("conn established")
	} else {
		logger.Info("conn store")
	}
	defer Hub.Delete(key)

	// 写消息
	go func() {
		for msg := range c.Write {
			err := conn.WriteMessage(msg.Type, msg.Data)
			if err != nil {
				logger.Error("write message error", zap.Error(err))
			} else {
				logger.Info("write message success", zap.Int("type", msg.Type), zap.String("data", string(msg.Data)))
			}
		}
	}()

	// 读消息
	for {
		t, p, e := conn.ReadMessage()
		if websocket.IsCloseError(e, websocket.CloseNormalClosure) {
			logger.Info("client normal close")
			break
		} else if e != nil {
			logger.Error("read message error", zap.Error(e))
			break
		}

		logger.Info("read message success", zap.Int("type", t), zap.String("data", string(p)))
		if f != nil {
			f(key, t, p)
		}
	}
	return nil
}
