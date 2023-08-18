package goquery

import "github.com/PuerkitoBio/goquery"

var (
	// 解析html文档
	NewDocumentFromReader = goquery.NewDocumentFromReader
)

type (
	// 节点集
	// Find() 查找节点
	// First() 选择第一个节点
	// Html() gets the HTML contents of the first element
	Selection = goquery.Selection
)
