# [pods](https://kubernetes.io/docs/concepts/workloads/pods/)

k8s通过pod来管理容器，并通过复制pod实现水平扩展，这样一组pod构成工作负载

pod是类似物理机或虚拟机的逻辑主机，是一组相对耦合的容器，共享网络和存储等资源

The "one-container-per-Pod" model is the most common Kubernetes use case
