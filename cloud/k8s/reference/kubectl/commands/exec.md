# [exec](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#exec)

Execute a command in a container.

```bash
# 执行单个命令
kubectl exec (POD | TYPE/NAME) [-c CONTAINER] [flags] -- COMMAND [args...]
# 进入容器终端
kubectl exec -it POD -- bash
```
