package main

import (
	"github.com/violetpupil/gos/lib/gin"
	_ "github.com/violetpupil/gos/lib/gin/examples/docs"
)

func main() {
	gin.Run(":8082")
}
