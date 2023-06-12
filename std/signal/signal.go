package signal

import (
	"os"
	"os/signal"
	"syscall"
)

var (
	// Stop causes package signal to stop relaying incoming signals to c.
	Stop = signal.Stop
)

// GracefullyExit 阻塞等待 INT TERM 信号
func GracefullyExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
