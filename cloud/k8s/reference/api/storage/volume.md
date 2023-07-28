# [Volume](https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/volume/)

- `name` 卷名，pod内唯一 required
- `configMap` 填充卷的ConfigMap
  - `name` ConfigMap名
  - `items` 投影对象键对应文件相对路径
- `emptyDir` 空卷
- `hostPath` 主机路径
  - `path` 路径

## KeyToPath

- `key` 投影对象的键
- `path` 相对于挂载路径的文件路径
