# [workloads](https://kubernetes.io/docs/concepts/workloads/)

workload是多个pod组成的应用

`Deployment` is a good fit for managing a stateless application workload on your cluster

`StatefulSet` lets you run one or more related Pods that do track state somehow.

`DaemonSet` defines Pods that provide node-local facilities.

`Job` 一次性任务

`CronJob` 重复执行的任务
