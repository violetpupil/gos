package main

import (
	"fmt"
	"net/http"

	"github.com/violetpupil/gos/lib/logrus"
	"github.com/violetpupil/gos/lib/ws"
)

func main() {
	logrus.DisableQuote()
	handler := new(ws.Websocket)
	fmt.Println("websocket listen on 0.0.0.0:8080")
	err := http.ListenAndServe(":8080", handler)
	fmt.Println(err)
}
