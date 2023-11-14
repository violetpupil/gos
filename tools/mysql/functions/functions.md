# [functions](https://dev.mysql.com/doc/refman/8.0/en/functions.html)

[Built-In Function and Operator Reference](https://dev.mysql.com/doc/refman/8.0/en/built-in-function-reference.html)

```sql
# 根据得分划分级别
SELECT elt( INTERVAL ( score, 0, 50, 100 ), "B", "A", "S" ) AS `level` FROM t
```
