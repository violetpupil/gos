# [cron](https://pkg.go.dev/github.com/robfig/cron/v3)

```text
Field name   | Allowed values  | Allowed special characters
----------   | --------------  | --------------------------
Minutes      | 0-59            | * / , -
Hours        | 0-23            | * / , -
Day of month | 1-31            | * / , - ?
Month        | 1-12 or JAN-DEC | * / , -
Day of week  | 0-6 or SUN-SAT  | * / , - ?
```

## Predefined schedules

```text
Entry                  | Description                                | Equivalent To
-----                  | -----------                                | -------------
@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 1 1 *
@monthly               | Run once a month, midnight, first of month | 0 0 1 * *
@weekly                | Run once a week, midnight between Sat/Sun  | 0 0 * * 0
@daily (or @midnight)  | Run once a day, midnight                   | 0 0 * * *
@hourly                | Run once an hour, beginning of hour        | 0 * * * *
```
