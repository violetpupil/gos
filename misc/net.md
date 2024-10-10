# [net](https://pkg.go.dev/net)

```golang
// 解析ip地址，返回nil则不是ip地址
net.ParseIP(ip)
// 解析域名或ip地址，返回一个ip地址
net.ResolveIPAddr("ip", address)
// 解析域名或ip地址，返回多个ip地址
net.LookupHost(host)
```
