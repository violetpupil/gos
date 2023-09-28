build:
	set GOOS=linux\
	&set CGO_ENABLED=0\
	&go build -o app .\
	&docker build -t instafever/resource-info:v1.0.0 .\
	&docker push instafever/resource-info:v1.0.0