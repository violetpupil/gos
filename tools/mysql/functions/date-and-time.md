# [date-and-time](https://dev.mysql.com/doc/refman/8.0/en/date-and-time-functions.html)

`CURDATE()` Return the current date

`CURTIME()` Return the current time

`DATE()` Extract the date part of a date or datetime expression

`FROM_UNIXTIME()` 将秒级时间戳转为时间字符串

`MONTH()` 获取月份

`NOW()` Return the current date and time

`TIME()` Extract the time portion of the expression passed

`UNIX_TIMESTAMP()` 返回秒级时间戳

```sql
# 指定时区
SELECT CONVERT_TZ(FROM_UNIXTIME(1711414800), 'UTC', 'Asia/Shanghai') AS datetime;
```

## format

`%Y` 四位年

`%m` 两位月

`%d` 两位日

`%H` 小时 00-23

`%i` 分钟

`%S` `%s` 秒
