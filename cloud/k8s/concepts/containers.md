# [containers](https://kubernetes.io/docs/concepts/containers/)

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

## [images](https://kubernetes.io/docs/concepts/containers/images/)

### Image names

`<image-name>:<tag>`

`<image-name>@<digest>`

You should avoid using the :latest tag when deploying containers in production as it is harder to track which version of the image is running and more difficult to roll back properly.

### imagePullPolicy

`IfNotPresent` 如果本地没有镜像，则拉取。设置镜像tag非latest时的默认值

`Always` 先向镜像注册表查询镜像摘要，如果本地没有，则拉取。镜像tag是latest时的默认值

`Never` 不拉取镜像，如果本地没有镜像，则启动失败

镜像拉取策略在对象创建时设置，修改tag后，不会自动更新

拉取镜像失败后，采用BackOff重试，最大重试间隔是5分钟

### Serial and parallel image pulls

The kubelet never pulls multiple images in parallel on behalf of one Pod.

However, if you have two Pods that use different images, the kubelet pulls the images in parallel on behalf of the two different Pods, when parallel image pulls is enabled.

## [Lifecycle Hooks](https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/)

`PostStart` 容器创建后执行

`PreStop` 容器停止前执行

If either a PostStart or PreStop hook fails, it kills the Container.

## [Container states](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-states)

`Waiting` 准备状态，拉取镜像

`Running` 运行中，PostStart钩子已执行

`Terminated` 已停止

## [Garbage collection of unused containers and images](https://kubernetes.io/docs/concepts/architecture/garbage-collection/#containers-images)

The kubelet performs garbage collection on unused images every five minutes and on unused containers every minute.

当磁盘用量达到 HighThresholdPercent 时，会删除最久未使用的镜像，直到磁盘用量降到 LowThresholdPercent

## [Container probes](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes)

A probe is a diagnostic performed periodically by the kubelet on a container.

`exec` 在容器中执行命令，退出码为0时成功

`grpc` grpc健康检查，响应status=SERVING时成功

`httpGet` http get请求，响应码在200~399时成功

`tcpSocket` tcp连接上就成功，可以立即关闭连接

### Types of probe

`livenessProbe` 容器是否在运行，失败会重启

`readinessProbe` 容器是否能处理请求，失败会把pod从服务端点中移除

`startupProbe` 容器是否已启动，失败会重启

All other probes are disabled if a startup probe is provided, until it succeeds.
