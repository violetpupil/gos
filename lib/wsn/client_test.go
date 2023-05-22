package wsn

import "testing"

func TestChat_Chat(t *testing.T) {
	c := new(Chat)
	c.Chat("ws://localhost:8080")
}
