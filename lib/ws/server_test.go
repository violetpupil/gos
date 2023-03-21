package ws

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWebsocket_ServeHTTP(t *testing.T) {
	handler := new(Websocket)
	err := http.ListenAndServe(":8080", handler)
	fmt.Println(err)
}
