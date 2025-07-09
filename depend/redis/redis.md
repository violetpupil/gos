# [redis](https://redis.io/)

```bash
docker run -dp 6379:6379 -v /root/redis:/data --name redis redis
```

管道不一定都执行

事务要么都执行，要么都不执行，但不保证都成功
