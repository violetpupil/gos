package nsq

import (
	"os"
	"os/signal"
	"syscall"

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

// Consume 消费消息，addr是nsqlookupd地址
func Consume(addr, topic, channel string, hs ...Handler) error {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		logrus.Errorln("new consumer error", err)
		return err
	}

	// 添加消息处理器
	for _, h := range hs {
		consumer.AddHandler(h)
	}
	err = consumer.ConnectToNSQLookupd(addr)
	if err != nil {
		logrus.Errorln("connect nsqlookupd error", err)
		return err
	}

	// 通过 INT TERM 信号停止消费
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	consumer.Stop()
	<-consumer.StopChan
	return nil
}
