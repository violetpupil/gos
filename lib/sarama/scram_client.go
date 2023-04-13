// https://github.com/Shopify/sarama/tree/main/examples/sasl_scram_client
package sarama

import (
	"github.com/violetpupil/components/lib/logrus"
	"github.com/xdg-go/scram"
)

// SCRAMClient sarama.SCRAMClient实现
type SCRAMClient struct {
	*scram.Client
	*scram.ClientConversation
	// 工厂函数，返回hash.Hash
	scram.HashGeneratorFcn
}

// Begin 生成客户端
func (x *SCRAMClient) Begin(userName, password, authzID string) (err error) {
	x.Client, err = x.HashGeneratorFcn.NewClient(userName, password, authzID)
	if err != nil {
		logrus.Errorln("new client error", err)
		return err
	}
	x.ClientConversation = x.Client.NewConversation()
	return nil
}

// Step 执行scram交换
func (x *SCRAMClient) Step(challenge string) (response string, err error) {
	response, err = x.ClientConversation.Step(challenge)
	return
}

// Done scram会话是否结束
func (x *SCRAMClient) Done() bool {
	return x.ClientConversation.Done()
}
