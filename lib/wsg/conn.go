package wsg

import (
	"sync"

	"github.com/gorilla/websocket"
)

var Hub sync.Map

type Conn struct {
	Conn *websocket.Conn
}
