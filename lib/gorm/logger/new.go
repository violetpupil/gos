package logger

import (
	"log"
	"os"
	"time"

	gLog "gorm.io/gorm/logger"
)

// Colors
const (
	Reset = "\033[0m"    // 重置
	R     = "\033[31m"   // 红
	G     = "\033[32m"   // 绿
	Y     = "\033[33m"   // 黄
	B     = "\033[34m"   // 蓝
	M     = "\033[35m"   // 紫
	C     = "\033[36m"   // 淡蓝
	W     = "\033[37m"   // 白
	BB    = "\033[34;1m" // 蓝加粗
	MB    = "\033[35;1m" // 紫加粗
	RB    = "\033[31;1m" // 红加粗
	YB    = "\033[33;1m" // 黄加粗
)

// Color 将文本上色，然后重置
func Color(c string, s string) string {
	return c + s + Reset
}

// OneLine 单行日志logger
func OneLine() gLog.Interface {
	writer := log.New(os.Stdout, "\r\n", log.LstdFlags)
	config := gLog.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  gLog.Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	}

	infoStr := Color(G, "%s") + Color(G, "[info] ")
	warnStr := Color(BB, "%s") + Color(M, "[warn] ")
	errStr := Color(M, "%s") + Color(R, "[error] ")
	traceStr := Color(G, "%s") + Y + "[%.3fms] " + Color(BB, "[rows:%v]") + " %s"
	traceWarnStr := G + "%s " + Color(Y, "%s") + RB + "[%.3fms] " + Y + "[rows:%v]" + Color(M, " %s")
	traceErrStr := RB + "%s " + Color(MB, "%s") + Y + "[%.3fms] " + Color(BB, "[rows:%v]") + " %s"

	return &Logger{
		Writer:       writer,
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}
