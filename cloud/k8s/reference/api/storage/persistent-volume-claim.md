# [PersistentVolumeClaim](https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/persistent-volume-claim-v1/)

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
spec:
status:
```

## PersistentVolumeClaimSpec

- `accessModes` 存取模式
- `resources` 资源用量 storage
  - `requests` describes the minimum amount of compute resources required
- `storageClassName` StorageClass名
