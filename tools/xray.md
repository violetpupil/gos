# [Xray](https://xtls.github.io/)

## [文档](https://xtls.github.io/document/)

```bash
# 安装
bash -c "$(curl -L https://github.com/XTLS/Xray-install/raw/main/install-release.sh)" @ install
# 创建配置
vi /usr/local/etc/xray/config.toml
# 修改服务配置
vi /etc/systemd/system/xray.service
ExecStart=/usr/local/bin/xray run -config /usr/local/etc/xray/config.toml
# 重启服务
systemctl restart xray
# 查看服务状态
systemctl status xray
# 验证
curl -x socks5://user:pass@127.0.0.1:2023 http://ipecho.net/plain
```

## [配置](https://xtls.github.io/config/)

```toml
# 将127.0.0.1换成代理ip
# 每个代理ip应该设置不同的tag
[[inbounds]]
listen = "127.0.0.1"
port = 2023
protocol = "socks"
tag = "1"

[inbounds.settings]
auth = "password"
udp = true
ip = "127.0.0.1"

[[inbounds.settings.accounts]]
user = "user"
pass = "pass"

[[routing.rules]]
type = "field"
inboundTag = "1"
outboundTag = "1"

[[outbounds]]
sendThrough = "127.0.0.1"
protocol = "freedom"
tag = "1"
```
