# 环境变量只在此次make生效
win:
	set GOOS=linux&set CGO_ENABLED=0&go build -o tmp .