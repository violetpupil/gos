# [cli](https://docs.docker.com/engine/reference/commandline/cli/)

`login` 登录注册处

`logout` 登出注册处

## [常用命令]

```bash
# 上传镜像到公共镜像库
docker push instafever/component
# 使用当前目录Dockerfile构建镜像
docker build -t instafever/component .
# Remove all dangling images.
docker image prune -f
```

## [build](https://docs.docker.com/engine/reference/commandline/build/)

`-t` Name and optionally a tag in the name:tag format

## [run](https://docs.docker.com/engine/reference/commandline/run/)

`--rm` Automatically remove the container when it exits

## [stop](https://docs.docker.com/engine/reference/commandline/stop/)

The main process inside the container will receive SIGTERM, and after a grace period, SIGKILL.
