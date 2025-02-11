# mysql

```bash
docker run -d --name mysql \
--privileged=true \
-p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=root \
-v ~/Documents/docker/mysql:/var/lib/mysql \
mysql
```

```sql
SELECT VERSION();
```
