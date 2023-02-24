package net

import (
	"errors"
	"net"

	"github.com/sirupsen/logrus"
)

// InterfaceUp 网卡是否启用
func InterfaceUp(i net.Interface) bool {
	return i.Flags&net.FlagUp != 0
}

// InterfaceBroadcast 网卡支持广播
func InterfaceBroadcast(i net.Interface) bool {
	return i.Flags&net.FlagBroadcast != 0
}

// InterfaceLoopback 回环网卡
func InterfaceLoopback(i net.Interface) bool {
	return i.Flags&net.FlagLoopback != 0
}

// InterfaceMulticast 网卡支持多播
func InterfaceMulticast(i net.Interface) bool {
	return i.Flags&net.FlagMulticast != 0
}

// InterfacePhysical 是否为物理网卡
func InterfacePhysical(i net.Interface) bool {
	if InterfaceLoopback(i) {
		return false
	}
	if i.Name == "vEthernet (WSL)" {
		return false
	}
	return true
}

// Interfaces 启用的物理网卡
func Interfaces() ([]net.Interface, error) {
	ifs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	dst := make([]net.Interface, 0)
	for _, i := range ifs {
		if InterfacePhysical(i) && InterfaceUp(i) {
			dst = append(dst, i)
		}
	}
	return dst, nil
}

// Ipv4 网卡ipv4地址
func Ipv4(i net.Interface) (net.IP, error) {
	addrs, err := i.Addrs()
	if err != nil {
		return nil, err
	}
	for _, a := range addrs {
		ipnet, ok := a.(*net.IPNet)
		// 是ipv4地址
		if ok && ipnet.IP.To4() != nil {
			return ipnet.IP, nil
		}
	}
	return nil, errors.New("the interface does not have an ipv4 address")
}

// InterfacesIpv4 启用的物理网卡ipv4地址
func InterfacesIpv4() ([]net.IP, error) {
	// 启用的物理网卡
	ifs, err := Interfaces()
	if err != nil {
		return nil, err
	}

	ips := make([]net.IP, 0)
	for _, i := range ifs {
		ip, err := Ipv4(i)
		if err != nil {
			logrus.Error("interface ipv4 error ", err)
			continue
		}
		ips = append(ips, ip)
	}
	return ips, nil
}
