# [prometheus](https://prometheus.io/)

[安装](https://prometheus.io/docs/prometheus/latest/installation/)

[文档](https://prometheus.io/docs/introduction/overview/)

[image](https://hub.docker.com/r/prom/prometheus)

[exporters](https://prometheus.io/docs/instrumenting/exporters/)

```bash
# prometheus.yml 要先创建
# http://localhost:9090/
docker run -dp 9090:9090 ^
-v d:/docker/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml ^
--name prometheus ^
prom/prometheus
```
