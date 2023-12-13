# [administration](https://dev.mysql.com/doc/refman/8.0/en/sql-server-administration-statements.html)

```sql
# 查看连接的客户端
show processlist
# 查看空闲连接最大空闲时长 秒
show variables like 'wait_timeout'
# 查看最大连接数
show variables like 'max_connections'
# 查看数据文件目录
SHOW VARIABLES LIKE 'datadir'
# 查看表索引
SHOW INDEX FROM tbl_name
# 关闭连接
kill CONNECTION processlist_id
```
