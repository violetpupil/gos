package signal

import (
	"os"
	"os/signal"
	"syscall"
)

var (
	// Stop causes package signal to stop relaying incoming signals to c.
	Stop = signal.Stop
	// Notify 将信号发到channel
	// If no signals are provided, all incoming signals will be relayed to c.
	// 可以将一个信号发给多个channel
	Notify = signal.Notify
)

// GracefullyExit 阻塞等待 INT TERM 信号
func GracefullyExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
