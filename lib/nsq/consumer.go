package nsq

import (
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

// Handler 实现 nsq.Handler 接口
type Handler func(*nsq.Message) error

// HandleMessage 处理消息
// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
func (h Handler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// A message with an empty body is simply ignored/discarded.
		return nil
	}

	logrus.WithFields(logrus.Fields{
		"ID":          string(m.ID[:]),
		"Body":        string(m.Body),
		"Timestamp":   m.Timestamp,
		"Attempts":    m.Attempts,
		"NSQDAddress": m.NSQDAddress,
	}).Info("handle message")
	err := h(m)
	return err
}
