# [EXPLAIN](https://dev.mysql.com/doc/refman/8.0/en/explain.html)

obtain a query execution plan

```sql
# type 输出格式 默认输出表格 TRADITIONAL 树型 TREE JSON
EXPLAIN [FORMAT=type] explainable_stmt
```

## [输出字段](https://dev.mysql.com/doc/refman/8.0/en/explain-output.html)

`key` 使用的索引

`type` ALL表示全表扫描
