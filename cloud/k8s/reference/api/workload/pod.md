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
- `image` 镜像
- `imagePullPolicy` Image pull policy.
- `command` 容器入口命令
- `ports` 公开端口列表
  - `containerPort` Number of port to expose on the pod's IP address. required
  - `name` 端口名，pod内唯一
- `env` List of environment variables to set in the container.
  - `name` 环境变量名
  - `value` 环境变量值
  - `valueFrom` Source for the environment variable's value.
    - `fieldRef` Selects a field of the pod
    - `secretKeyRef` 使用secret的一个键值
      - `key` secret中的键名
      - `name` secret名
- `volumeMounts` 挂载到容器的卷
  - `mountPath` 挂载路径 required
  - `name` 卷名 required
- `resources` 计算资源 cpu memory
  - `limits` Limits describes the maximum amount of compute resources allowed.
  - `requests` Requests describes the minimum amount of compute resources required.
- `lifecycle` 容器生命周期钩子 LifecycleHandler
  - `postStart` 容器创建后执行
  - `preStop` 容器正常停止前执行
- `readinessProbe` Periodic probe of container service readiness.
- `securityContext` 安全选项

## LifecycleHandler

- `exec` 执行的命令
  - `command` 单条命令，相对于/目录执行

## Probe

- `exec` 执行的命令
  - `command` 单条命令，相对于/目录执行
- `initialDelaySeconds` 容器启动到liveness探针启动间隔
- `timeoutSeconds` 探针超时秒数
