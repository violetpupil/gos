# [concepts](https://kubernetes.io/docs/concepts/)

## [Going back in time](https://kubernetes.io/docs/concepts/overview/)

![Deployment evolution](https://d33wubrfki0l68.cloudfront.net/26a177ede4d7b032362289c6fccd448fc4a91174/eb693/images/docs/container_evolution.svg)

### 物理机 vs 虚拟机 vs 容器

在物理机上，没法定义资源边界

虚拟机在虚拟硬件上，运行自己的操作系统

容器共享宿主机的操作系统

## [Addons](https://kubernetes.io/docs/concepts/overview/components/)

Because these are providing cluster-level features, namespaced resources for addons belong within the kube-system namespace.

集群自带 DNS, Dashboard, Container Resource Monitoring, Cluster-level Logging 插件
