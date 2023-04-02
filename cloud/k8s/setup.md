# [setup](https://kubernetes.io/docs/setup/)

## [container runtime](https://kubernetes.io/docs/setup/production-environment/container-runtimes/)

You need to install a container runtime into each node in the cluster so that Pods can run there.

Kubernetes 1.26 requires that you use a runtime that conforms with the Container Runtime Interface (CRI).

### linux control groups

It's critical that the kubelet and the container runtime uses the same cgroup driver and are configured the same.
