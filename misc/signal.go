package misc

import (
	"os"
	"os/signal"
	"syscall"
)

// GracefullyExit 阻塞等待 INT TERM 信号
// Ctrl+C发送SIGINT信号
// kill 命令默认发送 SIGTERM 信号
func GracefullyExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
