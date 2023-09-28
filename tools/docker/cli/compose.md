# [compose](https://docs.docker.com/compose/reference/)

```bash
# 创建并后台运行容器
docker compose up -d
# 实时看所有服务的日志
docker compose logs -f
# 移除网络和所有容器
docker compose down
```

## 选项

`-f` 指定compose file，默认./docker-compose.yml

`-p` 指定项目名，默认当前目录名

## [up](https://docs.docker.com/engine/reference/commandline/compose_up/)

`-d` 运行在后台

## logs 看所有服务的日志

`-f` 实时输出
