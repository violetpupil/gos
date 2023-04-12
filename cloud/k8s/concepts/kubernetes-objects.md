# [Kubernetes Objects](https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/)

Almost every Kubernetes object includes two nested object fields that govern the object's configuration: the object spec and the object status.

spec - desired state

status -  current state

## [unique](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/)

Every Kubernetes object also has a UID that is unique across your whole cluster.

Kubernetes UIDs are universally unique identifiers (also known as UUIDs).

## [labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)

用户定义数据

Equality-based

Set-based

## [namespaces](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)

有的对象是命名空间级的，有的是集群级的

k8s有四个初始命名空间 `default` `kube-node-lease` `kube-public` `kube-system`

Avoid creating namespaces with the prefix kube-, since it is reserved for Kubernetes system namespaces.
