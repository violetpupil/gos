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

## [Resources](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/#resources)

The Pod's effective request/limit for a resource is the higher of:

- the sum of all app containers request/limit for a resource
- the effective init request/limit for a resource

The highest of any particular resource request or limit defined on all init containers is the effective init request/limit.

## [Disruptions](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/)

主动执行的操作导致pod中断称为voluntary disruptions

A PDB limits the number of Pods of a replicated application that are down simultaneously from voluntary disruptions.

Cluster managers and hosting providers should use tools which respect PodDisruptionBudgets by calling the Eviction API instead of directly deleting pods or deployments.
