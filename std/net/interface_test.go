package net

import (
	"fmt"
	"testing"
)

func TestInterfaces(t *testing.T) {
	ifc, err := Interfaces()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", ifc)
}

func TestInterfacesIpv4Physical(t *testing.T) {
	fmt.Println(InterfacesIpv4Physical())
}

func TestInterfacesIpv4NL(t *testing.T) {
	fmt.Println(InterfacesIpv4NL())
}
