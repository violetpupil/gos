package net

import (
	"fmt"
	"net"
	"testing"
)

func TestInterfaces(t *testing.T) {
	ifc, err := Interfaces()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", ifc)
}

func TestInterfacesIpv4(t *testing.T) {
	fmt.Println(InterfacesIpv4())
}

func TestHostIp(t *testing.T) {
	fmt.Println(HostIp())
}

func TestPrivate(t *testing.T) {
	fmt.Println(Private(net.ParseIP("172.16.0.0")))
	fmt.Println(Private(net.ParseIP("172.32.0.0")))
}
