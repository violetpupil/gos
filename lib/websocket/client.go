package websocket

import (
	"context"

	"nhooyr.io/websocket"
)

// Chat 和websocket服务通信
func Chat(url string) {
	ctx := context.Background()
	websocket.Dial(ctx, url, nil)
}
