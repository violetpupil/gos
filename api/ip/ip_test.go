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
