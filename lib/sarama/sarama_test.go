package sarama

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama"
)

func TestSASLConfig(t *testing.T) {
	c := SASLConfig("", "")
	topics := []string{"my-topic"}
	err := Consume(
		"localhost:9092", "my-group", topics,
		func(cm *sarama.ConsumerMessage) {
			fmt.Printf("process %+v\n", cm)
		},
		c,
	)
	fmt.Println(err)
}
