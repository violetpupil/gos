package ws

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWebsocket_ServeHTTP(t *testing.T) {
	handler := new(Websocket)
	err := http.ListenAndServe("", handler)
	fmt.Println(err)
}
