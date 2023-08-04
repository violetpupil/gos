# [pod](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/)

```yaml
apiVersion: v1
kind: Pod
metadata:
spec:
status:
```

## PodSpec

- `containers` 容器列表 required
- `volumes` Volume列表
- `nodeSelector` node标签选择器
- `nodeName` is a request to schedule this pod onto a specific node.
- `affinity` affinity scheduling rules
  - `nodeAffinity` node affinity scheduling rules
- `tolerations` 可容忍的taints
  - `key` taint key
  - `operator` 默认Equal判断值相等，设为Exists判断键存在
  - `value` taint value
  - `effect` taint effect
- `schedulerName` 指定调度器
- `restartPolicy` 容器重启策略
- `terminationGracePeriodSeconds` term和kill信号间隔时间，默认30s

## Container

- `name` 容器名，pod内唯一 required
- `image` 镜像名
- `command` 容器入口命令
- `ports` 公开端口列表
  - `containerPort` Number of port to expose on the pod's IP address. required
  - `name` 端口名，pod内唯一
- `env` List of environment variables to set in the container.
  - `name` 环境变量名
  - `value` 环境变量值
- `volumeMounts` 挂载到容器的卷
  - `mountPath` 挂载路径 required
  - `name` 卷名 required
- `resources` 计算资源
  - `limits` Limits describes the maximum amount of compute resources allowed.
  - `requests` Requests describes the minimum amount of compute resources required.
