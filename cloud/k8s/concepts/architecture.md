# architecture

![Components of Kubernetes](https://d33wubrfki0l68.cloudfront.net/2475489eaf20163ec0f54ddc1d92aa8d4c87c96b/e7c81/images/docs/components-of-kubernetes.svg)

`etcd` 存储集群数据

`kube-scheduler` watches for newly created Pods with no assigned node, and selects a node for them to run on

`kubelet` 运行容器

## [controller](https://kubernetes.io/docs/concepts/architecture/controller/)

In Kubernetes, controllers are control loops that watch the state of your cluster.
Each controller tries to move the current cluster state closer to the desired state.

A controller tracks at least one Kubernetes resource type.

## [Addons](https://kubernetes.io/docs/concepts/overview/components/)

Because these are providing cluster-level features, namespaced resources for addons belong within the kube-system namespace.

集群自带 DNS, Dashboard, Container Resource Monitoring, Cluster-level Logging 插件
