# [nginx](https://nginx.org/)

[文档](https://nginx.org/en/docs/)

```bash
# 安装nginx
yum install -y nginx
# 运行nginx容器
# 必须先创建nginx.conf，否则会自动创建成文件夹
# html目录下index.html对应网站/
# https://hub.docker.com/_/nginx
docker run -dp 80:80 ^
-v D:/docker-volume/nginx/html:/usr/share/nginx/html:ro ^
-v D:/docker-volume/nginx/nginx.conf:/etc/nginx/nginx.conf:ro ^
--name nginx ^
nginx
```
