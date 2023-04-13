package nsq

import (
	"fmt"
	"testing"

	"github.com/nsqio/go-nsq"
)

func TestConsume(t *testing.T) {
	err := Consume(
		"localhost:4161",
		"myTopic",
		"myChannel",
		func(m *nsq.Message) error {
			fmt.Println(string(m.Body))
			return nil
		})
	fmt.Println(err)
}
