package net

import (
	"fmt"
	"net"
	"testing"
)

func TestHostIp(t *testing.T) {
	fmt.Println(HostIp())
}

func TestPrivate(t *testing.T) {
	fmt.Println(Private(net.ParseIP("172.16.0.0")))
	fmt.Println(Private(net.ParseIP("172.32.0.0")))
}
