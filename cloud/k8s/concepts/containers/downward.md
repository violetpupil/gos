# [Downward API](https://kubernetes.io/docs/concepts/workloads/pods/downward-api/)

The downward API allows containers to consume information about themselves or the cluster without using the Kubernetes client or API server.

## Pod-level fields

`metadata.name` the pod's name

`metadata.namespace` the pod's namespace

`metadata.uid` the pod's unique ID

`metadata.annotations['<KEY>']` the value of the pod's annotation

`metadata.labels['<KEY>']` the text value of the pod's label

### environment variables

`spec.serviceAccountName` the name of the pod's service account

`spec.nodeName` the name of the node where the Pod is executing

`status.hostIP` the primary IP address of the node to which the Pod is assigned

`status.podIP` the pod's primary IP address

### downwardAPI volume

`metadata.labels` all of the pod's labels

`metadata.annotations` all of the pod's annotations

## Container-level fields

`limits.cpu` A container's CPU limit

`requests.cpu` A container's CPU request

`limits.memory` A container's memory limit

`requests.memory` A container's memory request
