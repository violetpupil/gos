# 环境变量只在此次make生效
# 命令的开头要用tab，不能用空格
# 否则报错makefile:2: *** missing separator.  Stop.
win:
	set GOOS=linux\
	&set CGO_ENABLED=0\
	&go build -o tmp .