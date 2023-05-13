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
  - `name` All ports within a ServiceSpec must have unique names.
- `clusterIP` clusterIP is the IP address of the service and is usually assigned randomly.
