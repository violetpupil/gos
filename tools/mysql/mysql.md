# [mysql](https://www.mysql.com/)

[文档](https://dev.mysql.com/doc/refman/8.0/en/)

[tutorial](https://xiaolincoding.com/mysql/)

```bash
# 运行容器
docker run -d --name mysql ^
--privileged=true ^
-p 3306:3306 ^
-e MYSQL_ROOT_PASSWORD=root ^
-v d:/docker/mysql:/var/lib/mysql ^
mysql
```
