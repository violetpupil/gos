# [rollout](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#rollout)

Manage the rollout of a resource.

`history` View previous rollout revisions and configurations.

`pause` 暂停rollout

`resume` 继续rollout

`status` Show the status of the rollout.

`undo` Roll back to a previous rollout.

## history

CHANGE-CAUSE is copied from the Deployment annotation kubernetes.io/change-cause to its revisions upon creation.

`--revision` 查看指定版本细节

## undo

`--to-revision` The revision to rollback to. 默认是上一个版本
