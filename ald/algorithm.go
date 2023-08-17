package ald

import (
	"container/list"
	"sync"
)

// LRUList Least Recently Used列表
type LRUList struct {
	list.List
	sync.Mutex
}
