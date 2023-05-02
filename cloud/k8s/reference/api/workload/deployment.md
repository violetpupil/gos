# [Deployment](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/deployment-v1/)

```yaml
apiVersion: apps/v1
kind: Deployment
```

## DeploymentSpec

`selector` required 和pod模板标签一致的pod标签选择器

`template` required pod模板

`replicas` pod数，默认1
