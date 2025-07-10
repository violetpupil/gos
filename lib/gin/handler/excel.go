package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-the-way/exl"
	"github.com/sirupsen/logrus"
	"github.com/violetpupil/gos/std/mime"
)

type Fruit struct {
	// 没有excel标签的不处理
	ID   int    `excel:"ID"`
	Name string `excel:"Name"`
}

func (*Fruit) ReadConfigure(*exl.ReadConfig) {}

func (*Fruit) WriteConfigure(*exl.WriteConfig) {}

// Excel 生成excel
func Excel(c *gin.Context) {
	fruits := []*Fruit{{ID: 1, Name: "pear"}}
	c.Header("content-type", mime.Excel)
	// 自动写200响应码
	err := exl.WriteTo(c.Writer, fruits)
	if err != nil {
		logrus.Errorln("write to writer error", err)
	}
}
