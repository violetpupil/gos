package main

import "github.com/violetpupil/components/lib/websocket"

func main() {
	c := new(websocket.Chat)
	c.Chat("ws://localhost:8080")
}
