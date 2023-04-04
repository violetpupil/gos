# concepts

## [Going back in time](https://kubernetes.io/docs/concepts/overview/)

![Deployment evolution](https://d33wubrfki0l68.cloudfront.net/26a177ede4d7b032362289c6fccd448fc4a91174/eb693/images/docs/container_evolution.svg)

### 物理机 vs 虚拟机 vs 容器

在物理机上，没法定义资源边界

虚拟机在虚拟硬件上，运行自己的操作系统

容器共享宿主机的操作系统

## [Components](https://kubernetes.io/docs/concepts/overview/components/)

A Kubernetes cluster consists of a set of worker machines, called nodes, that run containerized applications.

![Components of Kubernetes](https://d33wubrfki0l68.cloudfront.net/2475489eaf20163ec0f54ddc1d92aa8d4c87c96b/e7c81/images/docs/components-of-kubernetes.svg)

### Control Plane

kube-apiserver

etcd 存储集群数据
