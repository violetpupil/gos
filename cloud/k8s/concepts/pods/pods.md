# [pods](https://kubernetes.io/docs/concepts/workloads/pods/)

k8s通过pod来管理容器，并通过复制pod实现水平扩展，这样一组pod构成工作负载

pod是类似物理机或虚拟机的逻辑主机，是一组相对耦合的容器，共享网络和存储资源

The "one-container-per-Pod" model is the most common Kubernetes use case

Controllers for workload resources create Pods from a pod template and manage those Pods on your behalf.

## Pod networking

Each Pod is assigned a unique IP address for each address family.

Inside a Pod (and only then), the containers that belong to the Pod can communicate with one another using localhost.

The containers in a Pod can also communicate with each other using standard inter-process communications like SystemV semaphores or POSIX shared memory.

Containers within the Pod see the system hostname as being the same as the configured name for the Pod.

## Static Pods

静态pod直接被kubelet管理，API server只能看到静态pod，但不能控制

## [Resources](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/#resources)

The Pod's effective request/limit for a resource is the higher of:

- the sum of all app containers request/limit for a resource
- the effective init request/limit for a resource

The highest of any particular resource request or limit defined on all init containers is the effective init request/limit.

## [Pod Quality of Service Classes](https://kubernetes.io/docs/concepts/workloads/pods/pod-qos/)

k8s根据容器资源请求和限制，将pod分为不同的quality of service (QoS) class

When a Node runs out of resources, Kubernetes will first evict BestEffort Pods running on that Node, followed by Burstable and finally Guaranteed Pods.
