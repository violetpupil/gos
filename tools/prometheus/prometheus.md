# [prometheus](https://prometheus.io/)

[安装](https://prometheus.io/docs/prometheus/latest/installation/)

[文档](https://prometheus.io/docs/introduction/overview/)

```bash
# prometheus.yml 要先创建
docker run -dp 9090:9090 ^
-v d:/docker/prometheus/prometheus.yml:/prometheus/prometheus.yml ^
--name prometheus ^
prom/prometheus
```
