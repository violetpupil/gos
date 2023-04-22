# [workloads](https://kubernetes.io/docs/concepts/workloads/)

workload是多个pod组成的应用

`Deployment` is a good fit for managing a stateless application workload on your cluster

`StatefulSet` lets you run one or more related Pods that do track state somehow.

`Job` 一次性任务

`CronJob` 重复执行的任务

## [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/)

`DaemonSet` defines Pods that provide node-local facilities.

Every time you add a node to your cluster that matches the specification in a DaemonSet, the control plane schedules a Pod for that DaemonSet onto the new node.
