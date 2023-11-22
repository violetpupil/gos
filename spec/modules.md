# [Modules](https://go.dev/ref/mod)

[教程](https://zhuanlan.zhihu.com/p/613649591)

## env

### GOPROXY 模块代理

默认镜像代理 `https://proxy.golang.org`

国内镜像代理 `https://goproxy.cn` `https://goproxy.io`

```bash
# 设置模块代理
go env -w GOPROXY=https://goproxy.cn,direct
# direct 回源到 vcs 地址抓取
# 可解决镜像同步不及时的问题
go env -w GOPROXY=direct
```
