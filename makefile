# https://www.ruanyifeng.com/blog/2015/02/make.html
# 环境变量只在此次make生效
# 默认执行第一个target
# 命令的开头要用tab，不能用空格
# 否则报错makefile:2: *** missing separator.  Stop.
linux:
	set GOOS=linux\
	&set CGO_ENABLED=0\
	&go build -o tmp .

darwin:
	set GOOS=darwin\
	&set CGO_ENABLED=0\
	&go build -o tmp-darwin .