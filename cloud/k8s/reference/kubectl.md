# [kubectl](https://kubernetes.io/docs/reference/kubectl/)

## [install](https://kubernetes.io/docs/tasks/tools/install-kubectl-windows/)

You must use a kubectl version that is within one minor version difference of your cluster.

配置文件 ~/.kube/config

## [Kubernetes Object Management](https://kubernetes.io/docs/concepts/overview/working-with-objects/object-management/)

kubectl有三种管理k8s对象的方式

1. 命令式命令 在命令中指定操作、对象配置

2. 命令式对象配置 在命令中指定操作、对象配置文件

3. 声明式对象配置 在命令中指定对象配置文件，操作由kubectl自动检测

## [commands](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands)

`get` 查看k8s资源

`delete` 删除k8s资源

### KUBECTL SETTINGS AND USAGE

`api-resources` 打印集群支持的api资源

`config` Modifies kubeconfig files.

## [常用命令](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)

```bash
kubectl cluster-info # 查看集群信息
kubectl get namespace # 查看所有命名空间
kubectl api-resources # 打印命名空间级的资源
kubectl api-resources --namespaced=false # 打印非命名空间级的资源
kubectl config view # Show Merged kubeconfig settings.
kubectl config set-context --current --namespace=np # 设置当前上下文的命名空间
kubectl version --client --output=yaml # 查看kubectl版本
```
