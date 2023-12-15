package main

import (
	"github.com/violetpupil/gos/lib/gin"
	_ "github.com/violetpupil/gos/lib/gin/examples/docs"
)

// @title gin示例
// @version v0.0.1
func main() {
	gin.Run(":8082")
}
