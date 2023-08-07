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
- `updateStrategy` pod更新策略
  - `type` 策略类型 Default is RollingUpdate.
  - `rollingUpdate` 滚动更新配置
    - `partition` pod序号小于该值不更新，大于等于该值更新，默认0
- `podManagementPolicy` 串行还是并行scale，默认串行
- `volumeClaimTemplates` PersistentVolumeClaim列表
- `minReadySeconds` pod启动并稳定运行几秒后，滚动更新下一个pod
