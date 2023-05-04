# [deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

## Creating a Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
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

## Updating a Deployment

A Deployment's rollout is triggered if and only if the Deployment's Pod template (that is, .spec.template) is changed

滚动更新确保75%~125%的pod运行，先创建部分pod，再杀掉部分pod
