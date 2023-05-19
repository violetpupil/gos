# [Job](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/job-v1/)

```yaml
apiVersion: batch/v1
kind: Job
metadata:
spec:
status:
```

## JobSpec

- `template` pod模板 required
- `parallelism` 并行运行的pod数，默认1
- `completions` 运行成功的pod数，默认1
- `completionMode` pod运行模式
- `backoffLimit` 重试次数，默认6
