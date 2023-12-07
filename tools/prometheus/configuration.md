# [configuration](https://prometheus.io/docs/prometheus/latest/configuration/configuration/)

```yaml
scrape_configs:
  - job_name: app
    scrape_interval: 10s
    static_configs:
      - targets:
          - localhost:2112
```
