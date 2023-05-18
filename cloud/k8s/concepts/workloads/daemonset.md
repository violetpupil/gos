# [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/)

Every time you add a node to your cluster that matches the specification in a DaemonSet, the control plane schedules a Pod for that DaemonSet onto the new node.

Kubernetes places DaemonSet Pods on nodes before they are ready

If you do not specify either, then the DaemonSet controller will create Pods on all nodes.

If node labels are changed, the DaemonSet will promptly add Pods to newly matching nodes and delete Pods from newly not-matching nodes.
