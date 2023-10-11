# [Aggregate Functions](https://dev.mysql.com/doc/refman/8.0/en/aggregate-functions-and-modifiers.html)

## GROUP_CONCAT()

连接非 NULL 值成字符串，没有非 NULL 值，返回 NULL

```sql
# 连接常量 -> 12
SELECT GROUP_CONCAT('a', 'b')
# 连接查询字段 -> 1,2
SELECT GROUP_CONCAT(id) FROM t_account
```
