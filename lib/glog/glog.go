// flag.Parse must be called before any logging is done.
package glog

import "github.com/golang/glog"

var (
	Errorln = glog.Errorln
	// 记录fatal日志，然后退出程序
	Exitln = glog.Exitln
	// 记录fatal日志，然后退出程序
	// including a stack trace of all running goroutines
	Fatalln   = glog.Fatalln
	Infoln    = glog.Infoln
	Warningln = glog.Warningln
)
