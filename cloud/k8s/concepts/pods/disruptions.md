# [Disruptions](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/)

主动执行的操作导致pod中断称为voluntary disruptions

## Pod disruption budgets

A PDB limits the number of Pods of a replicated application that are down simultaneously from voluntary disruptions.

Cluster managers and hosting providers should use tools which respect PodDisruptionBudgets by calling the Eviction API instead of directly deleting pods or deployments.

## DisruptionTarget condition reason

`PreemptionByScheduler` 被更高优先级的pod挤掉

`DeletionByTaintManager` 被Taint Manager删除

`EvictionByEvictionAPI` 由于驱逐api调用被驱逐

`DeletionByPodGC` 节点不再存在，被pod垃圾回收删除

`TerminationByKubelet` 因节点资源压力或关机被终结
