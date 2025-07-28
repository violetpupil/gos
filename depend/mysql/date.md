# date

```sql
# 当前时间戳
UNIX_TIMESTAMP(NOW())
# 今天0点时间戳
UNIX_TIMESTAMP(CURDATE())
# 一天前时间戳
UNIX_TIMESTAMP(DATE_SUB(NOW(), INTERVAL 1 DAY))
# 前一天0点时间戳
UNIX_TIMESTAMP(DATE_SUB(CURDATE(), INTERVAL 1 DAY))
# 时间戳转时间
from_unixtime(1742381794)
# 0时区转8时区
CONVERT_TZ(Now(), '+00:00', '+08:00')
```
