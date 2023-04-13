package nsq

import (
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

// Producer will lazily connect to that instance (and re-connect)
// when Publish commands are executed.
var Producer *nsq.Producer

// InitProducer 初始化生产者
func InitProducer(addr string) error {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(addr, config)
	if err != nil {
		logrus.Errorln("new producer error", err)
		return err
	}
	Producer = producer
	return nil
}

// StopProducer 停止生产者
func StopProducer() {
	Producer.Stop()
}

// Publish 同步发送单个消息
func Publish(topic string, body []byte) error {
	err := Producer.Publish(topic, body)
	return err
}

// PublishStr 同步发送单个消息字符串
func PublishStr(topic, body string) error {
	return Publish(topic, []byte(body))
}
