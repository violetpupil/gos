package sarama

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// ConsumerGroupHandler 处理单个分区
// 实现sarama.ConsumerGroupHandler接口，方法会被并发调用
type ConsumerGroupHandler struct {
	// 消费单条消息
	Consume func(*sarama.ConsumerMessage)
}

// Setup 处理之前初始化
func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

// Cleanup 处理结束清理
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim 处理消息
func (h ConsumerGroupHandler) ConsumeClaim(s sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		logrus.WithFields(logrus.Fields{
			"Timestamp":      msg.Timestamp,
			"BlockTimestamp": msg.BlockTimestamp,
			"Key":            string(msg.Key),
			"Value":          string(msg.Value),
			"Topic":          msg.Topic,
			"Partition":      msg.Partition,
			"Offset":         msg.Offset,
		}).Infoln("consume message")
		h.Consume(msg)
		s.MarkMessage(msg, "")
	}
	return nil
}

// Consume 使用消费者组消费，f为处理消息函数
func Consume(addr, groupId, topic string, f func(*sarama.ConsumerMessage)) error {
	config := sarama.NewConfig()
	addrs := []string{addr}
	group, err := sarama.NewConsumerGroup(addrs, groupId, config)
	if err != nil {
		logrus.Errorln("new consumer group error", err)
		return err
	}
	defer func() {
		err := group.Close()
		logrus.WithField("Error", err).Infoln("close group")
	}()

	for {
		ctx := context.Background()
		topics := []string{topic}
		handler := ConsumerGroupHandler{f}

		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			logrus.Errorln("consume error", err)
			return err
		}
	}
}
