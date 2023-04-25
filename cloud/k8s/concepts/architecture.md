# architecture

![Components of Kubernetes](https://d33wubrfki0l68.cloudfront.net/2475489eaf20163ec0f54ddc1d92aa8d4c87c96b/e7c81/images/docs/components-of-kubernetes.svg)

`etcd` 存储集群数据

`kube-scheduler` watches for newly created Pods with no assigned node, and selects a node for them to run on

`kubelet` 运行容器

## [nodes](https://kubernetes.io/docs/concepts/architecture/nodes/)

节点会向集群发送心跳，有两种形式：

1. 更新节点的status字段

2. Each Node has an associated Lease object.

## [controller](https://kubernetes.io/docs/concepts/architecture/controller/)

In Kubernetes, controllers are control loops that watch the state of your cluster.
Each controller tries to move the current cluster state closer to the desired state.

A controller tracks at least one Kubernetes resource type.

## [garbage-collection](https://kubernetes.io/docs/concepts/architecture/garbage-collection/)

### Cascading deletion

主对象删除后，k8s默认会删除从对象

默认使用后台删除，先删除主对象，再删除从对象

可以指定前台删除，先删除从对象，再删除主对象

### Garbage collection of unused containers and images

The kubelet performs garbage collection on unused images every five minutes and on unused containers every minute.

当磁盘用量达到 HighThresholdPercent 时，会删除最久未使用的镜像，直到磁盘用量降到 LowThresholdPercent
