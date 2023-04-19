# [setup](https://kubernetes.io/docs/setup/)

## [container runtime](https://kubernetes.io/docs/setup/production-environment/container-runtimes/)

You need to install a container runtime into each node in the cluster so that Pods can run there.

The Kubernetes [Container Runtime Interface (CRI)](https://kubernetes.io/docs/concepts/architecture/cri/) defines the main gRPC protocol for the communication between the cluster components kubelet and container runtime.

### linux control groups

It's critical that the kubelet and the container runtime uses the same cgroup driver and are configured the same.

```bash
# 查看操作系统使用的cgroup版本
# v2 -> cgroup2fs
# v1 -> tmpfs
stat -fc %T /sys/fs/cgroup/
```
