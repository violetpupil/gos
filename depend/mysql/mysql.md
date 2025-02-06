# mysql

```bash
docker run -d --name mysql ^
--privileged=true ^
-p 3306:3306 ^
-e MYSQL_ROOT_PASSWORD=root ^
-v d:/docker/mysql:/var/lib/mysql ^
mysql
```

## 常见问题

### 1. select查询字符串匹配，大小写不敏感

```sql
select 'a' = 'A'
select cast('a' as binary) = 'A'
```

### 2. 分页排序，相同值乱序

按一个无索引字段排序，相同的值排序是不固定的。分页的时候，会出现相同的数据出现在两页上。
