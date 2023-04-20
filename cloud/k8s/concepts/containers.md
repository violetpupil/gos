# [containers](https://kubernetes.io/docs/concepts/containers/)

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
