# [kubectl](https://kubernetes.io/docs/reference/kubectl/)

For configuration, kubectl looks for a file named config in the $HOME/.kube directory.

## Syntax

`kubectl [command] [TYPE] [NAME] [flags]`

Resource types are case-insensitive and you can specify the singular, plural, or abbreviated forms.

## [install](https://kubernetes.io/docs/tasks/tools/install-kubectl-windows/)

You must use a kubectl version that is within one minor version difference of your cluster.

## [Kubernetes Object Management](https://kubernetes.io/docs/concepts/overview/working-with-objects/object-management/)

kubectl有三种管理k8s对象的方式

1. 命令式命令 在命令中指定操作、对象配置

2. 命令式对象配置 在命令中指定操作、对象配置文件

3. 声明式对象配置 在命令中指定对象配置文件，操作由kubectl自动检测

## Resource types

events ev

namespaces ns

nodes no

persistentvolumeclaims pvc

pods po

services svc

deployments deploy

replicasets rs

statefulsets sts

horizontalpodautoscalers hpa
