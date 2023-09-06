package runtime

import (
	"runtime"
)

const (
	OSWindows = "windows"
	OSDarwin  = "darwin"
	OSLinux   = "linux"
)

var (
	// To view possible combinations of GOOS and GOARCH, run "go tool dist list".
	// 当前cpu架构
	GOARCH = runtime.GOARCH
	// 当前操作系统
	GOOS = runtime.GOOS
)
