package net

import "net"

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
