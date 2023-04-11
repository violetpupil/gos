# [SigNoz](https://signoz.io/)

[文档](https://signoz.io/docs/)

[instrumentation](https://signoz.io/docs/instrumentation/)

## [install](https://signoz.io/docs/install/docker/)

SigNoz can be installed on macOS or Linux computers

```bash
git clone -b main https://github.com/SigNoz/signoz.git && cd signoz/deploy/
docker-compose -f docker/clickhouse-setup/docker-compose.yaml up -d
```

[dashboard](http://localhost:3301/)
