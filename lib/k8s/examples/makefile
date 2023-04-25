build:
	set GOOS=linux&go build -o ./app .
	docker build -t instafever/component .
	del app
	docker push instafever/component