# [Kubernetes Objects](https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/)

Almost every Kubernetes object includes two nested object fields that govern the object's configuration: the object spec and the object status.

spec - desired state

status -  current state

## [unique](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/)

Every Kubernetes object also has a UID that is unique across your whole cluster.

Kubernetes UIDs are universally unique identifiers (also known as UUIDs).

## [namespaces](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)

k8s有四个初始命名空间 `default` `kube-node-lease` `kube-public` `kube-system`

Avoid creating namespaces with the prefix kube-, since it is reserved for Kubernetes system namespaces.

## [labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)&[annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/)

labels 用户定义标识

annotations 用户定义非标识数据

label和annotation键名可以包含前缀，用/分隔

The kubernetes.io/ and k8s.io/ prefixes are reserved for Kubernetes core components.

[Shared labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/common-labels/) and annotations share a common prefix: app.kubernetes.io.

## [owners](https://kubernetes.io/docs/concepts/overview/working-with-objects/owners-dependents/)

In Kubernetes, some objects are owners of other objects.
These owned objects are dependents of their owner.

Cluster-scoped dependents can only specify cluster-scoped owners.

命名空间级的从对象能指定集群级或同一命名空间的主对象

## [Cascading deletion](https://kubernetes.io/docs/concepts/architecture/garbage-collection/#cascading-deletion)

主对象删除后，k8s默认会删除从对象

默认使用后台删除，先删除主对象，再删除从对象

可以指定前台删除，先删除从对象，再删除主对象
