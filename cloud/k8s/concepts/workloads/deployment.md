# [deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

## Creating a Deployment

```bash
kubectl apply -f https://k8s.io/examples/controllers/nginx-deployment.yaml
```

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

A Deployment's revision history is stored in the ReplicaSets it controls.

Deployment --是主对象--> ReplicaSet --是主对象--> Pod

## Updating a Deployment

A Deployment's rollout is triggered if and only if the Deployment's Pod template (that is, .spec.template) is changed

滚动更新确保75%~125%的pod运行，maxUnavailable向下取整，maxSurge向上取整
