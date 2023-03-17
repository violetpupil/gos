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

func TestInterfacesIpv4(t *testing.T) {
	fmt.Println(InterfacesIpv4())
}
