# [StatefulSets](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/)

StatefulSets currently require a Headless Service to be responsible for the network identity of the Pods.

StatefulSet的pod名固定为`statefulset名-ordinal`，ordinal从0开始

StatefulSet给pod添加`statefulset.kubernetes.io/pod-name`标签
