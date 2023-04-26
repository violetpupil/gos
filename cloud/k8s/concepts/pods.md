# [pods](https://kubernetes.io/docs/concepts/workloads/pods/)

k8s通过pod来管理容器，并通过复制pod实现水平扩展，这样一组pod构成工作负载

pod是类似物理机或虚拟机的逻辑主机，是一组相对耦合的容器，共享网络和存储资源

The "one-container-per-Pod" model is the most common Kubernetes use case

## Pod templates

Controllers for workload resources create Pods from a pod template and manage those Pods on your behalf.

模板更新后，需要创建新的pod替换旧的

## Pod networking

Each Pod is assigned a unique IP address for each address family.

Inside a Pod (and only then), the containers that belong to the Pod can communicate with one another using localhost.

The containers in a Pod can also communicate with each other using standard inter-process communications like SystemV semaphores or POSIX shared memory.

Containers within the Pod see the system hostname as being the same as the configured name for the Pod.

## [Pod Lifecycle](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/)

Pods are only scheduled once in their lifetime.

### Pod phase

`Pending` pod还没有调度到node上，或者容器还在准备

`Running` pod被调度到node上，所有容器被创建，至少一个容器处于启动或运行中

`Succeeded` All containers in the Pod have terminated in success, and will not be restarted.

`Failed` All containers in the Pod have terminated, and at least one container has terminated in failure.

`Unknown` 与节点通信有问题，无法获取pod状态

### Container restart policy

The spec of a Pod has a restartPolicy field with possible values Always, OnFailure, and Never.
The default value is Always.

容器退出后，kubelet采用指数回退重启，最多五分钟。
容器正常运行10分钟后，重置回退计时器
