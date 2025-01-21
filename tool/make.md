# make

<https://www.ruanyifeng.com/blog/2015/02/make.html>

命令的开头要用tab，不能用空格

否则报错makefile:2: *** missing separator.  Stop.

```makefile
linux:
 set GOOS=linux\
 &set CGO_ENABLED=0\
 &go build -o app-linux .

darwin:
 set GOOS=darwin\
 &set CGO_ENABLED=0\
 &go build -o app-darwin .
```
