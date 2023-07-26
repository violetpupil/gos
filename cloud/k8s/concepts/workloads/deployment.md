# [deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

## Creating a Deployment

```bash
# 用文件创建nginx deployment
kubectl apply -f https://k8s.io/examples/controllers/nginx-deployment.yaml
# 用命令创建nginx deployment
kubectl create deployment nginx --image=nginx:1.14.2
# 创建service，暴露nginx deployment
kubectl expose deployment nginx --type=LoadBalancer --port=80
# 本地端口映射
kubectl port-forward service/nginx 8080:80
# 查看日志
kubectl logs -f deploy/nginx
# 进入终端
kubectl exec -it deploy/nginx -- bash
# 更新镜像
kubectl set image deployments/nginx nginx=nginx:1.25.1
```

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: nginx
  name: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```

Deployment controller将`pod-template-hash`标签添加到 ReplicaSet 和 Pod 上

A Deployment's revision history is stored in the ReplicaSets it controls.

Deployment --是主对象--> ReplicaSet --是主对象--> Pod

## Updating a Deployment

A Deployment's rollout is triggered if and only if the Deployment's Pod template (that is, .spec.template) is changed

滚动更新确保75%~125%的pod运行，maxUnavailable向下取整，maxSurge向上取整
