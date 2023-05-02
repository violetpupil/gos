# [Pod Quality of Service Classes](https://kubernetes.io/docs/concepts/workloads/pods/pod-qos/)

k8s根据容器资源请求和限制，将pod分为不同的quality of service (QoS) class

When a Node runs out of resources, Kubernetes will first evict BestEffort Pods running on that Node, followed by Burstable and finally Guaranteed Pods.

## Guaranteed

They are guaranteed not to be killed until they exceed their limits or there are no lower-priority Pods that can be preempted from the Node.

They may not acquire resources beyond their specified limits.

- Every Container in the Pod must have a memory limit and a memory request.
- For every Container in the Pod, the memory limit must equal the memory request.
- Every Container in the Pod must have a CPU limit and a CPU request.
- For every Container in the Pod, the CPU limit must equal the CPU request.

## Burstable

Because a Burstable Pod can include a Container that has no resource limits or requests, a Pod that is Burstable can try to use any amount of node resources.

- At least one Container in the Pod has a memory or CPU request or limit.

## BestEffort

Pods in the BestEffort QoS class can use node resources that aren't specifically assigned to Pods in other QoS classes.
