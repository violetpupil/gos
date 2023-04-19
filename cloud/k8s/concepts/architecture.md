# architecture

![Components of Kubernetes](https://d33wubrfki0l68.cloudfront.net/2475489eaf20163ec0f54ddc1d92aa8d4c87c96b/e7c81/images/docs/components-of-kubernetes.svg)

## [nodes](https://kubernetes.io/docs/concepts/architecture/nodes/)

节点会向集群发送心跳，有两种形式：

1. 更新节点的status字段

2. Each Node has an associated Lease object.

## [controller](https://kubernetes.io/docs/concepts/architecture/controller/)

In Kubernetes, controllers are control loops that watch the state of your cluster.
Each controller tries to move the current cluster state closer to the desired state.

A controller tracks at least one Kubernetes resource type.

## [garbage-collection](https://kubernetes.io/docs/concepts/architecture/garbage-collection/)

By default, Kubernetes uses background cascading deletion unless you manually use foreground deletion or choose to orphan the dependent objects.

The kubelet performs garbage collection on unused images every five minutes and on unused containers every minute.
