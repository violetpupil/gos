# [Deployment](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/deployment-v1/)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
spec:
status:
```

## DeploymentSpec

`selector` 和pod模板标签一致的pod标签选择器 required

`template` pod模板 required

`replicas` pod数，默认1

`revisionHistoryLimit` The number of old ReplicaSets to retain to allow rollback. 默认10

## DeploymentStatus

`conditions` deployment状况

- `status` Status of the condition, one of True, False, Unknown.
- `type` 状况类型
- `reason` The reason for the condition's last transition.
