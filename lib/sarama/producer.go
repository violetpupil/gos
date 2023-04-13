package sarama

import (
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// Produce 循环5s生产消息，topic会自动创建
// config如果传nil，会自动创建配置
func Produce(addr, topic, msg string, config *sarama.Config) error {
	if config == nil {
		config = sarama.NewConfig()
	}
	config.Producer.Return.Successes = true
	addrs := []string{addr}
	// 异步生产
	producer, err := sarama.NewAsyncProducer(addrs, config)
	if err != nil {
		logrus.Errorln("new async producer error", err)
		return err
	}

	// 消费生产结果channel
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for m := range producer.Successes() {
			logrus.Infof("produce success %+v", m)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			logrus.Errorln("produce error", err)
		}
	}()

	// 使用 ctrl-c 发送 SIGINT 信号停止
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

ProducerLoop:
	for {
		value := sarama.StringEncoder(msg)
		message := &sarama.ProducerMessage{Topic: topic, Value: value}
		select {
		case <-time.After(5 * time.Second):
			producer.Input() <- message
		case <-signals:
			producer.AsyncClose()
			break ProducerLoop
		}
	}

	// 等待生产结果channel消费完
	wg.Wait()
	return nil
}
