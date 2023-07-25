# [3proxy](https://github.com/3proxy/3proxy)

## windows

[下载](https://github.com/3proxy/3proxy/releases/download/0.9.4/3proxy-0.9.4-x64.zip)

1. 在bin64目录配置文件3proxy.cfg
2. 在bin64目录直接运行3proxy.exe

## linux

```bash
# 下载安装包
wget https://github.com/3proxy/3proxy/releases/download/0.9.4/3proxy-0.9.4.x86_64.rpm
# 安装
yum install -y ./3proxy-0.9.4.x86_64.rpm
# 卸载
yum remove 3proxy
# 重启服务
systemctl restart 3proxy
# 查看服务状态
systemctl status 3proxy
# 设置开机自启
systemctl enable --now 3proxy
# 创建pid文件
mkdir -p /var/run/3proxy
touch /var/run/3proxy/3proxy.pid
# 创建配置
cd /usr/local/3proxy/conf
mv 3proxy.cfg 3proxy.cfg.bk
vi 3proxy.cfg
# 创建日志目录
mkdir /usr/local/3proxy/logs
chmod -R 777 /usr/local/3proxy/logs
# 测试代理
curl -x socks5://M3aHRm2U0HteJJdJ:AG97Rzi6yKXE@localhost:1080 http://ipecho.net/plain
```

## 配置

```conf
# 日志文件，相对于/usr/local/3proxy目录
log /logs/3proxy-%y%m%d.log D
# 设置用户 user:CL:password
users M3aHRm2U0HteJJdJ:CL:AG97Rzi6yKXE
# 验证用户
auth strong
# socks端口
socks -p1080
```
