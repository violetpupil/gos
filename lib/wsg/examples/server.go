package main

import (
	"fmt"
	"net/http"

	"github.com/violetpupil/gos/lib/wsg"
)

func main() {
	http.HandleFunc("/ws", wsg.Upgrade)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
