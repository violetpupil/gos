# [DaemonSet](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/daemon-set-v1/)

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
spec:
status:
```

## DaemonSetSpec

- `selector` 和pod模板标签一致的pod标签选择器 required
- `template` pod模板 required
