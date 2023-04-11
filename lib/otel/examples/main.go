package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/violetpupil/components/lib/otel"
)

func main() {
	closeExporter, err := otel.InitTracer("", "")
	if err != nil {
		panic(err)
	}
	defer closeExporter(context.Background())

	e := gin.Default()
	otel.GinMiddleware(e)
	fmt.Println(e.Run())
}
