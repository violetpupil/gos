package ip

import (
	"fmt"
	"testing"
)

func TestProxyIpEcho(t *testing.T) {
	ip, err := ProxyIpEcho(
		"http://proxyserver:8888",
		"user",
		"pass",
	)
	fmt.Println(ip, err)
}

func TestProxyIpEchoFive(t *testing.T) {
	ip, err := ProxyIpEchoFive(
		"http",
		"host",
		"port",
		"user",
		"pass",
	)
	fmt.Println(ip, err)
}

func TestProxyIpLum(t *testing.T) {
	ip, err := ProxyIpLum(
		"http://proxyserver:8888",
		"user",
		"pass",
	)
	fmt.Println(ip, err)
}

func TestProxyIpLumFive(t *testing.T) {
	ip, err := ProxyIpLumFive(
		"http",
		"host",
		"port",
		"user",
		"pass",
	)
	fmt.Println(ip, err)
}

var proxyS = []string{}

func TestProxyIpEchoStr(t *testing.T) {
	for _, proxy := range proxyS {
		ip, err := ProxyIpEchoStr(proxy)
		fmt.Println(ip, err)
	}
}
