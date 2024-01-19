// 删除操作
// https://gorm.io/docs/delete.html
package gorm

import (
	"gorm.io/gorm"
)

type delete struct {
	db *gorm.DB
}

// Delete 删除数据，value是数据模型指针或模型切片指针
// 当value主键有值
// DELETE from emails where id = 10;
// var email = Email{ID: 10}
// db.Delete(&email)
//
// 当主键是数字
// DELETE FROM users WHERE id = 10;
// db.Delete(&User{}, 10)
// db.Delete(&User{}, "10") 小心sql注入
// DELETE FROM users WHERE id IN (1,2,3);
// db.Delete(&users, []int{1,2,3})
func (d *delete) Delete(value interface{}, cond ...interface{}) (int64, error) {
	db := d.db.Delete(value, cond...)
	return db.RowsAffected, db.Error
}
