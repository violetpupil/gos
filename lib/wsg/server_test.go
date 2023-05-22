package wsg

import (
	"fmt"
	"net/http"
	"testing"
)

func TestUpgrade(t *testing.T) {
	http.HandleFunc("/ws", Upgrade)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
