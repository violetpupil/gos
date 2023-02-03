package main

import (
	"fmt"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// main 运行http服务在端口8080，直到错误发生
func main() {
	e := gin.Default()

	// 注册pprof处理器，提供运行时数据，路径以/debug/pprof/开头
	pprof.Register(e)

	fmt.Print(e.Run())
}
