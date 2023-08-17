package list

import "container/list"

var (
	// 创建双向链表
	New = list.New
)

type (
	// 双向链表
	// Back() 最后一个元素
	// Front() 第一个元素
	// InsertAfter() 插入到mark元素后
	// InsertBefore() 插入到mark元素前
	// Len() 长度
	// MoveAfter() 移动到mark元素后
	// MoveBefore() 移动到mark元素前
	// MoveToBack() 移动到尾
	// MoveToFront() 移动到头
	// PushBack() 插入到尾
	// PushFront() 插入到头
	// Remove() 移除元素
	List = list.List

	// 双向链表元素
	// Value 元素值
	// Next() 下一个元素
	// Prev() 前一个元素
	Element = list.Element
)
