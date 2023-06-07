# [redis](https://redis.io/)

[文档](https://redis.io/docs/)

[命令](https://redis.io/commands/)

```bash
docker run -dp 6379:6379 -v /root/redis:/data --name redis redis
```

## [管道](https://redis.io/docs/manual/pipelining/)

不一定都执行

## [事务](https://redis.io/docs/manual/transactions/)

要么都执行，要么都不执行，但不保证都成功
