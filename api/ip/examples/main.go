package main

import "github.com/violetpupil/components/api/ip"

func main() {
	ip.ProxyInterval(
		"http://proxyserver:8888",
		"user",
		"pass",
	)
}
