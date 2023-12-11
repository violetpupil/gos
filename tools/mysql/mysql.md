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

## [查询流程](https://xiaolincoding.com/mysql/base/how_select.html)

连接器 -- 解析器 -- 预处理器 -- 优化器 -- 执行器 -- 存储引擎

## [Indexes](https://xiaolincoding.com/mysql/index/index_interview.html)

[为什么 MySQL 采用 B+ 树作为索引？](https://xiaolincoding.com/mysql/index/why_index_chose_bpuls_tree.html)

## 常见问题

[MySQL排序分页查询数据顺序错乱的原因和解决办法](https://blog.csdn.net/weixin_44299027/article/details/121627609)
