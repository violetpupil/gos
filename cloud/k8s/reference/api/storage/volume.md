# [Volume](https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/volume/)

- `name` 卷名，pod内唯一 required
- `configMap` 填充卷的ConfigMap，data键是文件，data值是内容
  - `name` ConfigMap名
  - `items` data键不是文件时指定文件
- `emptyDir` 空卷
- `hostPath` 主机路径
  - `path` 路径

## KeyToPath

- `key` 投影对象的键
- `path` 相对于挂载路径的文件路径
