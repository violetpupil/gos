# [containers](https://kubernetes.io/docs/concepts/containers/)

## Image names

`<image-name>:<tag>`

`<image-name>@<digest>`

## imagePullPolicy

`IfNotPresent` 如果本地没有镜像，则拉取

`Always` 先向镜像注册表查询镜像摘要，如果本地没有，则拉取

`Never` 不拉取镜像，如果本地没有镜像，则启动失败

You should avoid using the :latest tag when deploying containers in production as it is harder to track which version of the image is running and more difficult to roll back properly.
