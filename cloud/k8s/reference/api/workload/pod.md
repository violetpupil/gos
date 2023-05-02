# [pod](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/)

```yaml
apiVersion: v1
kind: Pod
metadata:
spec:
status:
```

## PodSpec

### Containers

`containers` 容器列表 required

### Lifecycle

`restartPolicy` 容器重启策略

## Container

`name` 容器名，pod内唯一 required

### Image

`image` 镜像名

### Ports

`ports` 公开端口列表

- `containerPort` Number of port to expose on the pod's IP address. required

### Resources

`resources` 计算资源

- `limits` Limits describes the maximum amount of compute resources allowed.
- `requests` Requests describes the minimum amount of compute resources required.
