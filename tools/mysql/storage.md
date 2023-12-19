# storage

[MySQL 一行记录是怎么存储的？](https://xiaolincoding.com/mysql/base/row_format.html)

`.frm` 表结构

`.ibd` 表数据，分为共享表空间文件和独占表空间文件

## 表空间文件

表空间由段（segment）、区（extent）、页（page）、行（row）组成

InnoDB 的数据是按「页」为单位来读写的

默认每个页的大小为 16KB
