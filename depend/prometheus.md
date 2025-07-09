# [prometheus](https://prometheus.io/)

```bash
# prometheus.yml 要先创建
# http://localhost:9090/
docker run -dp 9090:9090 ^
-v d:/docker/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml ^
--name prometheus ^
prom/prometheus
```
