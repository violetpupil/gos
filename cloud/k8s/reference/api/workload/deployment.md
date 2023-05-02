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
