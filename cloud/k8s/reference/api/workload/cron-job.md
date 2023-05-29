# [CronJob](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/cron-job-v1/)

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
spec:
status:
```

## CronJobSpec

- `jobTemplate` job模板 required
  - `metadata` 元数据
  - `spec` job spec
- `schedule` The schedule in Cron format. required
