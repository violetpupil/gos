package sarama

import (
	"fmt"
	"testing"
)

func TestProduce(t *testing.T) {
	err := Produce("localhost:9092", "my-topic", "hi", nil)
	fmt.Println(err)
}
