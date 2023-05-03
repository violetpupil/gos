# [常用命令](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)

```bash
kubectl get namespace # 查看所有命名空间
kubectl version --client --output=yaml # 查看kubectl版本
```

## Kubectl context and configuration

```bash
kubectl config view # Show Merged kubeconfig settings.
kubectl config set-context --current --namespace=<namespace-name> # 设置当前上下文的命名空间
```

## Viewing and finding resources

```bash
kubectl describe node <node-name> # 查看节点信息
```

## Updating resources

```bash
kubectl set image deployment/frontend www=image:v2 # 更新deployment的www镜像
```

## Interacting with Nodes and cluster

```bash
kubectl cordon <node-name> # Mark node as unschedulable.
kubectl cluster-info # 查看集群信息
kubectl api-resources # 打印命名空间级的资源
kubectl api-resources --namespaced=false # 打印非命名空间级的资源
```
