# select

## Locking Reads

在事务中，查询并加行锁，防止之后插入或更新数据时，查询行被其他事务修改。

如果查询没有使用索引，MySQL无法确定需要锁定哪些行，因此会`锁表`。

`FOR SHARE` 共享锁，多个事务可以同时持有，其他事务可以读，不能修改

操作同一行会造成死锁

- 事务 A：

```sql
START TRANSACTION;
SELECT * FROM table_name WHERE id = 1 FOR SHARE;
```

- 事务 B：

```sql
START TRANSACTION;
SELECT * FROM table_name WHERE id = 1 FOR SHARE; -- 允许
UPDATE table_name SET column_name = 'value' WHERE id = 1; -- 阻塞，直到事务 A 释放锁
```

`FOR UPDATE` 排他锁，只有一个事务可以持有，其他事务不能读取或修改

- 事务 A：

```sql
START TRANSACTION;
SELECT * FROM table_name WHERE id = 1 FOR UPDATE;
```

- 事务 B：

```sql
START TRANSACTION;
SELECT * FROM table_name WHERE id = 1 FOR SHARE; -- 阻塞
SELECT * FROM table_name WHERE id = 1 FOR UPDATE; -- 阻塞
UPDATE table_name SET column_name = 'value' WHERE id = 1; -- 阻塞
```
