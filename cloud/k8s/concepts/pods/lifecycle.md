# [Pod Lifecycle](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/)

Pods are only scheduled once in their lifetime.

## Pod phase

`Pending` pod还没有调度到node上，或者容器还在准备

`Running` pod被调度到node上，所有容器被创建，至少一个容器处于启动或运行中

`Succeeded` All containers in the Pod have terminated in success, and will not be restarted.

`Failed` All containers in the Pod have terminated, and at least one container has terminated in failure.

`Unknown` 与节点通信有问题，无法获取pod状态

## Container restart policy

The spec of a Pod has a restartPolicy field with possible values Always, OnFailure, and Never.
The default value is Always.

容器退出后，kubelet采用指数回退重启，最多五分钟。
容器正常运行10分钟后，重置回退计时器

## Pod conditions

`PodScheduled` pod是否已被调度到节点上

`PodHasNetwork` pod网络是否已配置

`ContainersReady` 是否所有容器已就绪

`Initialized` 是否所有初始容器已成功完成

`Ready` pod是否已可以处理请求

可以注入额外的状况，如果在pod状况列表中找不到该状况，则该状况状态默认为False

## Garbage collection of Pods

当pod数量超过配置值时，pod垃圾回收控制器(PodGC)会清理终结的pod

当节点不再存在，orphan pods会被回收

## [Pod restart reasons](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/#pod-restart-reasons)

A Pod can restart, causing re-execution of init containers:

- The Pod infrastructure container is restarted.
- 所有容器终结，初始容器完成记录已被清理
