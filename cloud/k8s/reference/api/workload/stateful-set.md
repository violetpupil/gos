# [StatefulSet](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/stateful-set-v1/)

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
spec:
status:
```

## StatefulSetSpec

- `serviceName` is the name of the service that governs this StatefulSet. required
- `selector` 和pod模板标签一致的pod标签选择器 required
- `template` pod模板 required
- `replicas` pod数，默认1
- `volumeClaimTemplates` PersistentVolumeClaim列表
