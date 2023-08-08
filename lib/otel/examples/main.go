package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/violetpupil/gos/lib/otel"
)

func main() {
	closeExporter, err := otel.InitTracer("localhost:4317", "goGinApp")
	if err != nil {
		panic(err)
	}
	defer closeExporter(context.Background())

	e := gin.Default()
	e.Use(otel.GinMiddleware())
	fmt.Println(e.Run(":8081"))
}
