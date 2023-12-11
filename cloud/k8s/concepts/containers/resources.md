# [resources](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/)

request 调度器用它决定pod运行在哪个node上，kubelet用它保留资源供容器使用

limit kubelet用它限制容器最大使用资源

如果指定了limit，没有指定request，request会使用limit值

`cpu` 1代表1核，指定小数限制cpu时间，0.1或者100m，最小1m

`memory` 指定字节数，支持单位G, M, k和Gi, Mi, Ki，单位m表示小于1字节
