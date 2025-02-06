# mysql

`order by`默认升序

```bash
docker run -d --name mysql ^
--privileged=true ^
-p 3306:3306 ^
-e MYSQL_ROOT_PASSWORD=root ^
-v d:/docker/mysql:/var/lib/mysql ^
mysql
```

## 常见问题

1. select查询字符串匹配，大小写不敏感

```sql
select 'a' = 'A'
select cast('a' as binary) = 'A'
```
