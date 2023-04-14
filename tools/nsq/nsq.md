# [nsq](https://nsq.io/)

[文档](https://nsq.io/overview/design.html)

## [docker](https://nsq.io/deployment/docker.html)

```bash
docker run --name nsqlookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd
# host使用宿主机ip
docker run --name nsqd \
-p 4150:4150 \
-p 4151:4151 \
nsqio/nsq /nsqd \
--broadcast-address=<host> \
--lookupd-tcp-address=<host>:4160 \
--data-path=/data
# 测试nsqlookupd
curl http://localhost:4161/ping
```
