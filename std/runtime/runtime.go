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

var (
	// 阻塞运行垃圾回收
	GC = runtime.GC
	// 填充runtime.MemStats实例
	ReadMemStats = runtime.ReadMemStats
)

type (
	// 内存分配统计
	// Mallocs is the cumulative count of heap objects allocated.
	// Frees is the cumulative count of heap objects freed.
	// The number of live objects is Mallocs - Frees.
	MemStats = runtime.MemStats
)
