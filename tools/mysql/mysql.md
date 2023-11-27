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

```sql
# 匹配字符串时，大小写不敏感
select 'A' = 'a'
```

## 常见问题

[MySQL排序分页查询数据顺序错乱的原因和解决办法](https://blog.csdn.net/weixin_44299027/article/details/121627609)
