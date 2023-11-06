# [nginx](https://nginx.org/)

[文档](https://nginx.org/en/docs/)

```bash
# 安装nginx
yum install -y nginx
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

## 配置

[ngx_http_core_module](https://nginx.org/en/docs/http/ngx_http_core_module.html)

[ngx_http_ssl_module](https://nginx.org/en/docs/http/ngx_http_ssl_module.html)

```conf
# https 配置
http {
    # 访问 http 全都重定向到 https
    server {
        listen 80;
        server_name _;
        rewrite ^(.*) https://domain$1 permanent;
    }

    # https 访问
    # 打开防火墙端口
    # 将证书放到 /etc/nginx/ca 目录下
    server {
        listen 443 ssl;
        server_name _;
        ssl_certificate /etc/nginx/ca/domain.pem;
        ssl_certificate_key /etc/nginx/ca/domain.key;
        ssl_session_timeout 5m;
        ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;
        ssl_protocols TLSv1.1 TLSv1.2;
        ssl_prefer_server_ciphers on;
    }
}
```
