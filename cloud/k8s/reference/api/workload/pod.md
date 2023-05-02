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

`containers` required 容器列表

### Lifecycle

`restartPolicy` 容器重启策略

## Container

`name` required 容器名，pod内唯一

### Image

`image` 镜像名

### Ports

`ports` 公开端口列表

- `containerPort` required pod公开端口号
