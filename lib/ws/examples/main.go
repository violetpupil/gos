package main

import (
	"fmt"
	"net/http"

	"github.com/violetpupil/components/lib/ws"
)

func main() {
	handler := new(ws.Websocket)
	err := http.ListenAndServe(":8080", handler)
	fmt.Println(err)
}
