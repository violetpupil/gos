# 终端工具

```bash
# 查询域名ip地址
nslookup www.baidu.com
# 设置 每次执行命令前，先打印
set -x
# 查看网络流量
iftop
```

```bash
# report file system disk space usage
# -h 使用单位
df -h
# 查看当前目录下文件夹大小
# -h 使用单位
# -r 降序
du -h --max-depth=1 | sort -hr
# 包括文件
du -ah --max-depth=1 | sort -hr
# 查看根目录，排除 /proc 路径
du -h --max-depth=1 --exclude /proc / | sort -hr
```
