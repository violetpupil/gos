# specialized containers

## [Pause container](https://kubernetes.io/docs/concepts/windows/intro/#pause-container)

[source code](https://github.com/kubernetes/kubernetes/tree/master/build/pause)

In Linux, the cgroups and namespaces that make up a pod need a process to maintain their continued existence;
the pause process provides this.

Kubernetes uses pause containers to allow for worker containers crashing or restarting without losing any of the networking configuration.

## [Init Containers](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/)

specialized containers that run before app containers in a Pod.

Each init container must complete successfully before the next one starts.

Consequently, they can be given access to Secrets that app containers cannot access.
