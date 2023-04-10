# [kubectl](https://kubectl.docs.kubernetes.io/)

## [install](https://kubernetes.io/docs/tasks/tools/install-kubectl-windows/)

You must use a kubectl version that is within one minor version difference of your cluster.

## [Kubernetes Object Management](https://kubernetes.io/docs/concepts/overview/working-with-objects/object-management/)

kubectl有三种管理k8s对象的方式

1. 命令式命令 在命令中指定操作、对象配置

2. 命令式对象配置 在命令中指定操作、对象配置文件

3. 声明式对象配置 在命令中指定对象配置文件，操作由kubectl自动检测

## [子命令](https://kubectl.docs.kubernetes.io/references/kubectl/)

delete 删除k8s资源

get 查看k8s资源

## 常用命令

```bash
kubectl get namespace # 查看所有命名空间
kubectl version --client --output=yaml # 查看kubectl版本
```
