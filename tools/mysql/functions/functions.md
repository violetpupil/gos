# [functions](https://dev.mysql.com/doc/refman/8.0/en/functions.html)

[Built-In Function and Operator Reference](https://dev.mysql.com/doc/refman/8.0/en/built-in-function-reference.html)

```sql
# 根据得分划分级别
SELECT elt( INTERVAL ( score, 0, 50, 100 ), "0-49", "50-99", "100-" ) AS `level` FROM t
# 根据得分划分级别 分组计数
SELECT elt( INTERVAL ( score, 0, 50, 100 ), "0-49", "50-99", "100-" ) AS `level`,
count(*) as count FROM t group by `level` order by `level`
```
