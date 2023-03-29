package sarama

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama"
)

func TestConsume(t *testing.T) {
	err := Consume(
		"localhost:9092", "my-group", "my-topic",
		func(cm *sarama.ConsumerMessage) {
			fmt.Printf("process %+v\n", cm)
		},
	)
	fmt.Println(err)
}
