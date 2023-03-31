package sarama

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama"
)

func TestConsume(t *testing.T) {
	topics := []string{"my-topic"}
	err := Consume(
		"localhost:9092", "my-group", topics,
		func(cm *sarama.ConsumerMessage) {
			fmt.Printf("process %+v\n", cm)
		},
	)
	fmt.Println(err)
}
