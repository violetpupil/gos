# [HorizontalPodAutoscaler](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/horizontal-pod-autoscaler-v2/)

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
spec:
status:
```

## HorizontalPodAutoscalerSpec

- `maxReplicas` 最大pod数 required
- `scaleTargetRef` reference to scaled resource. required
  - `kind` 种类
  - `name` 命名
  - `apiVersion` api版本
- `minReplicas` 最小pod数 默认1
