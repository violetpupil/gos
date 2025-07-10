# [nginx](https://nginx.org/)

```bash
# 运行nginx容器
# 必须先创建nginx.conf，否则会自动创建成文件夹
# html目录对应网站/，不指定路径则加载index.html
# https://hub.docker.com/_/nginx
docker run -dp 80:80 ^
-v D:/docker-volume/nginx/html:/usr/share/nginx/html:ro ^
-v D:/docker-volume/nginx/nginx.conf:/etc/nginx/nginx.conf:ro ^
--name nginx ^
nginx
```
