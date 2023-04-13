package nsq

import (
	"fmt"
	"testing"
)

func TestPublishStr(t *testing.T) {
	err := InitProducer("127.0.0.1:4150")
	if err != nil {
		panic(err)
	}
	err = PublishStr("myTopic", "hi")
	fmt.Println(err)
}
