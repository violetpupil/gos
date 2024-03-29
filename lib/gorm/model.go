// https://gorm.io/docs/models.html
// By default, GORM uses ID as primary key,
// pluralizes struct name to snake_cases as table name, snake_case as column name,
// and uses CreatedAt, UpdatedAt to track creating/updating time
// 使用TableName()方法自定义表名
//
// 模型tag `gorm:""`
// 键值用:连接，字段用;连接
// column 列名
// type column data type
// serializer 序列器 json/gob/unixtime
// primaryKey 主键
// unique 唯一
// index 索引
// comment 注释
//
// type
// 指定int type:int(32)
// 指定bigint type:bigint
// 默认字段类型
// int8 tinyint
// int bigint
// uint bigint
// string longtext
// bool tinyint
// datatypes.JSON json
// time.Time datetime
// gorm.DeletedAt datetime
package gorm

import "gorm.io/gorm"

type (
	// 基础模型
	Model = gorm.Model
)
