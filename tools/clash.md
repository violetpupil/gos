# [clash](https://github.com/Dreamacro/clash)

[文档](https://dreamacro.github.io/clash/)

```bash
# 创建配置目录
mkdir -p ~/.config/clash && cd ~/.config/clash
# 下载依赖
sudo yum install -y curl wget jq
# 下载配置文件
wget https://v1.socketproep.com/api/clash/trojan -O config.yaml
# 下载GeoIP数据库
wget https://github.com/Dreamacro/maxmind-geoip/releases/download/20230812/Country.mmdb
# 下载clash
wget https://github.com/Dreamacro/clash/releases/download/v1.17.0/clash-linux-amd64-v3-v1.17.0.gz
# 解压并删除压缩包
gzip -d clash*
# 重命名
mv clash* clash
# 添加可执行权限
chmod +x clash
# 移到全局可执行目录
sudo mv clash /usr/local/bin
# 修改配置
vi config.yaml
# 运行clash
nohup clash &> clash.log &
# 验证
curl -x socks5://user1:pass1@127.0.0.1:7890 http://ipecho.net/plain
```

## [config](https://dreamacro.github.io/clash/configuration/configuration-reference.html)

```yaml
# authentication of local SOCKS5/HTTP(S) server
authentication:
  - "user1:pass1"
# 允许局域网连接
allow-lan: true
# 代理配置
proxies:
  - name: us
    type: socks5
    server: 127.0.0.1
    port: 443
    username: username
    password: password
```
