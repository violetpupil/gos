# [man-pages](https://man7.org/linux/man-pages/)

[index](https://man7.org/linux/man-pages/dir_all_by_section.html)

1. user commands
2. system calls
3. library functions
4. special files
5. file formats and filesystems
6. games
7. overview and miscellany section
8. administration and privileged commands

## 常用命令

```bash
# 后台运行程序
# 标准输出和标准错误重定向到文件
# 追加写入文件
nohup ./server &>> server.log &
# 筛选进程id，并发送term信号
ps -ef | grep ./server | grep -v grep | awk '{print $2}' | xargs kill
```

## 常用文件

`/etc/sysconfig/network-script/ifcfg-eth0` 网卡配置
