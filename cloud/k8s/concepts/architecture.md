# architecture

## [nodes](https://kubernetes.io/docs/concepts/architecture/nodes/)

节点会向集群发送心跳，有两种形式：

1. 更新节点的status字段

2. Each Node has an associated Lease object.
