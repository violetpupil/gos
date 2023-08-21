package ald

import (
	"errors"
	"sync"
)

var ErrEmpty = errors.New("queue is empty")

// Queue 队列 线程安全
type Queue[T any] struct {
	Slice []T
	Lock  sync.Mutex
}

// Pop 弹出元素 先进先出
// 队列为空返回错误ErrEmpty
func (q *Queue[T]) Pop() (T, error) {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	var e T
	if len(q.Slice) < 1 {
		return e, ErrEmpty
	}
	e = q.Slice[0]
	q.Slice = q.Slice[1:]
	return e, nil
}
