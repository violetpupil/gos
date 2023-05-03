# [rollout](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#rollout)

`history` View previous rollout revisions and configurations.

`status` Show the status of the rollout.

`undo` Roll back to a previous rollout.

## history

CHANGE-CAUSE is copied from the Deployment annotation kubernetes.io/change-cause to its revisions upon creation.

`--revision` 查看指定版本细节

## undo

`--to-revision` The revision to rollback to. 默认是上一个版本
