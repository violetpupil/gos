# [set](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#set)

Configure application resources.

`image` Update existing container image(s) of resources.

`resources` 设置计算资源

## image

`kubectl set image (-f FILENAME | TYPE NAME) CONTAINER_NAME_1=CONTAINER_IMAGE_1`

## resources

For each compute resource, if a limit is specified and a request is omitted, the request will default to the limit.

`-c` 容器名，默认*表示资源下所有容器

`--limits` 资源限制
