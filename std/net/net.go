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

// Ipv4 网卡ipv4地址，使用4字节表示
func Ipv4(i net.Interface) ([]net.IP, error) {
	addrs, err := i.Addrs()
	if err != nil {
		return nil, err
	}

	ips := make([]net.IP, 0)
	for _, a := range addrs {
		ipnet, ok := a.(*net.IPNet)
		if ok {
			ip := ipnet.IP.To4()
			if ip != nil {
				ips = append(ips, ip)
			}
		}
	}
	return ips, nil
}

// InterfacesIpv4 启用的物理网卡ipv4地址，使用4字节表示
func InterfacesIpv4() ([]net.IP, error) {
	ifs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ips := make([]net.IP, 0)
	for _, i := range ifs {
		// 启用的物理网卡
		if InterfacePhysical(i) && InterfaceUp(i) {
			ipsIf, err := Ipv4(i)
			if err != nil {
				logrus.Error("interface ipv4 error ", err)
				continue
			}
			ips = append(ips, ipsIf...)
		}
	}
	return ips, nil
}

// HostIp 获取主机ip
// 8.8.8.8 和 8.8.4.4 是谷歌提供的dns服务器
func HostIp() (net.IP, error) {
	conn, err := net.Dial("ip:icmp", "8.8.8.8")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	addr, ok := conn.LocalAddr().(*net.IPAddr)
	if !ok {
		return nil, errors.New("not ip addr")
	}
	return addr.IP, nil
}

// Private 如果是私有ip，转为4字节表示返回，否则返回nil
// 私有ip地址范围
// 10.0.0.0 ~ 10.255.255.255
// 172.16.0.0 ~ 172.31.255.255
// 192.168.0.0 ~ 192.168.255.255
func Private(ip net.IP) net.IP {
	ip = ip.To4()
	if ip == nil {
		return nil
	}

	if ip[0] == 10 ||
		ip[0] == 172 && ip[1] >= 16 && ip[1] < 32 ||
		ip[0] == 192 && ip[1] == 168 {
		return ip
	}
	return nil
}
