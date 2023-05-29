# [CronJob](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/)

A CronJob creates Jobs on a repeating schedule.

## Concurrency policy

`Allow` 允许同时运行新旧job

`Forbid` 旧job还没有结束，跳过新job

`Replace` 用新job代替旧job
