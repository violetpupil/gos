// 查询操作
// https://gorm.io/docs/query.html
// https://gorm.io/docs/advanced_query.html
package gorm

import (
	"errors"

	"gorm.io/gorm"
)

// RecordNotFound 检查查询函数错误是否为 ErrRecordNotFound
func RecordNotFound(e error) bool {
	return errors.Is(e, gorm.ErrRecordNotFound)
}

type query struct {
	db *gorm.DB
}

// First 查找根据主键排序的第一条数据，无主键则根据第一个字段排序
// dest是数据模型指针，返回bool表示是否找到
// SELECT * FROM users ORDER BY id LIMIT 1;
//
// 当dest主键有值
// SELECT * FROM users WHERE id = 10;
// var user = User{ID: 10}
// db.First(&user)
//
// 当主键是数字
// SELECT * FROM users WHERE id = 10;
// db.First(&user, 10)
// db.First(&user, "10") 小心sql注入
//
// 主键不限定类型
// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";
// db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
func (q *query) First(dest interface{}, cond ...interface{}) (bool, error) {
	err := q.db.First(dest, cond...).Error
	if RecordNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Last 查找根据主键排序的最后一条数据，无主键则根据第一个字段排序
// dest是数据模型指针，返回bool表示是否找到
// SELECT * FROM users ORDER BY id DESC LIMIT 1;
func (q *query) Last(dest interface{}, cond ...interface{}) (bool, error) {
	err := q.db.Last(dest, cond...).Error
	if RecordNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Take 查找第一条数据
// dest是数据模型指针，返回bool表示是否找到
// SELECT * FROM users LIMIT 1;
func (q *query) Take(dest interface{}, cond ...interface{}) (bool, error) {
	err := q.db.Take(dest, cond...).Error
	if RecordNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Find 查询数据
// dest是数据模型切片指针，返回找到的记录数
// SELECT * FROM users;
//
// 当主键是数字
// SELECT * FROM users WHERE id IN (1,2,3);
// db.Find(&users, []int{1,2,3})
func (q *query) Find(dest interface{}, cond ...interface{}) (int64, error) {
	db := q.db.Find(dest, cond...)
	return db.RowsAffected, db.Error
}

// FirstOrCreate 查询第一条匹配记录，查不到则创建并赋值
// query为查询条件，attrs为创建时的附加字段
// assign为附加字段，查到则更新，查不到则创建时使用
// query、attrs、assign必须是数据模型或map
// dest是数据模型指针，返回插入或更新的记录数
//
// db.FirstOrCreate(&user, User{Name: "non_existing"})
// db.Where(User{Name: "non_existing"}).FirstOrCreate(&user)
// INSERT INTO "users" (name) VALUES ("non_existing");
func (q *query) FirstOrCreate(dest, query, attrs, assign any) (int64, error) {
	db := q.db.Where(query).Attrs(attrs).Assign(assign).FirstOrCreate(dest)
	return db.RowsAffected, db.Error
}
