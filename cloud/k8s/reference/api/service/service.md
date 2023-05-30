# [Service](https://kubernetes.io/docs/reference/kubernetes-api/service-resources/service-v1/)

```yaml
apiVersion: v1
kind: Service
metadata:
spec:
status:
```

## ServiceSpec

- `selector` Route service traffic to pods with label keys and values matching this selector.
- `ports` The list of ports that are exposed by this service.
  - `port` 端口号 required
  - `targetPort` pod端口，默认和port相等
  - `protocol` The IP protocol for this port. Default is TCP.
  - `name` All ports within a ServiceSpec must have unique names.
- `clusterIP` clusterIP is the IP address of the service and is usually assigned randomly.
