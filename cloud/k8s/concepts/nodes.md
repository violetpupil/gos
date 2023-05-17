# [nodes](https://kubernetes.io/docs/concepts/architecture/nodes/)

## Conditions

`Ready` True if the node is healthy and ready to accept pods

## Heartbeats

节点会向集群发送心跳，有两种形式：

1. 更新节点的status字段
2. Each Node has an associated Lease object.
