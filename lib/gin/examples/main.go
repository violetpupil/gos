package main

import (
	"github.com/violetpupil/components/lib/gin"
	_ "github.com/violetpupil/components/lib/gin/examples/docs"
)

func main() {
	gin.Run(":8082")
}
