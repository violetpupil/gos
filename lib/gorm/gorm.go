package gorm

import (
	"gorm.io/gorm"
)

// crud 数据库操作
// 必须先调用 Init 初始化
type crud struct {
	C *create
	R *query
	U *update
	D *delete
}

var Crud *crud

func Init(db *gorm.DB) {
	Crud = new(crud)
	Crud.C = &create{db}
	Crud.R = &query{db}
	Crud.U = &update{db}
	Crud.D = &delete{db}
}
