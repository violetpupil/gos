# [grafana](https://grafana.com/)

添加数据源，docker 部署 ip 不能用 localhost

添加仪表盘

```bash
# http://localhost:3000/
# 用户名密码都是 admin
docker run -dp 3000:3000 ^
--name=grafana grafana/grafana
```

## [dashboards](https://grafana.com/grafana/dashboards/)

[Go Runtime](https://grafana.com/grafana/dashboards/14061-go-runtime-metrics/)
