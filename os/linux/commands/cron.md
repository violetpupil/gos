# [cron](https://en.wikipedia.org/wiki/Cron)

[web editor](https://crontab.guru/)

```cron
# ┌───────────── minute (0 - 59)
# │ ┌───────────── hour (0 - 23)
# │ │ ┌───────────── day of the month (1 - 31)
# │ │ │ ┌───────────── month (1 - 12)
# │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday;
# │ │ │ │ │                                   7 is also Sunday on some systems)
# │ │ │ │ │
# │ │ │ │ │
# * * * * * <command to execute>
```

## macros

`@yearly` `@monthly` `@weekly` `@daily` `@hourly`

## crontab

`-e` edit user's crontab

`-l` list user's crontab
