# [nginx](https://nginx.org/)

[文档](https://nginx.org/en/docs/)

```bash
# 安装nginx
yum install -y nginx
# 运行nginx容器
docker run -d ^
-v D:/docker-volume/nginx:/usr/share/nginx/html:ro ^
--name nginx ^
nginx
```
