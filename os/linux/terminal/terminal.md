# 终端工具

```bash
# 后台运行程序
# &> 标准输出和标准错误重定向到文件
# >> 追加写入文件
# 程序不能操作server.log
nohup ./server &>> server.log &
# 筛选进程id，并发送term信号
ps -ef | grep ./server | grep -v grep | awk '{print $2}' | xargs kill
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
