// https://github.com/gobwas/ws-examples/blob/feature/cli/src/chat/client/main.go
package ws

import "github.com/gobwas/ws"

// 结构体
type (
	Dialer = ws.Dialer
)

// 变量
var (
	DefaultDialer = ws.DefaultDialer
)

// 函数
var (
	// Dial 连接到指定地址，并升级到websocket
	// 第二个返回值 bf 表示握手成功后服务器立即发送的消息
	// 如果没有立即发送，则为 nil
	// 当 bf 不为 nil 时，最好调用 ws.PutReader 将 reader 放回池里
	Dial = ws.Dial
)

// Chat 和websocket服务通信
func Chat(url string) {}
