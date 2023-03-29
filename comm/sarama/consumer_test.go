package sarama

import (
	"fmt"
	"testing"
)

func TestConsume(t *testing.T) {
	err := Consume("localhost:9092", "my-group", "my-topic")
	fmt.Println(err)
}
