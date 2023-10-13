# [nohup](https://www.ruanyifeng.com/blog/2016/02/linux-daemon.html)

退出终端后，在当前终端运行的进程可能收到SIGHUP信号

nohup阻止SIGHUP信号发给进程

但程序如果处理了SIGHUP信号，会覆盖nohup的阻止
