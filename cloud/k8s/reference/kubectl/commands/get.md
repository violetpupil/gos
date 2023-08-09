# [get](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#get)

查看k8s资源

```bash
# 列出所有工作负载和服务资源
kubectl get all
```

`-o` 输出格式 json|yaml|wide

`--show-labels` 显示labels

`-l` 标签选择器

`-w` After listing/getting the requested object, watch for changes.
