package sync

import "sync"

type (
	// Map 并发使用的map
	// CompareAndDelete() 值相等则删除
	// CompareAndSwap() 值相等则交换新值
	// Delete() 删除条目
	// Load() 返回值
	// LoadAndDelete() 删除条目并返回值
	// Range() 遍历条目
	// Store() 存储或更新条目
	// Swap() 交换值
	// LoadOrStore() 键存在直接返回值，否则先存储再返回
	// 返回 bool 表示是否直接 load
	Map = sync.Map

	// Mutex 互斥锁
	// Lock() 阻塞加锁
	// Unlock() 解锁
	Mutex = sync.Mutex

	// A WaitGroup waits for a collection of goroutines to finish.
	// Add() 增加计数
	// Done() decrements the WaitGroup counter by one.
	// Wait() blocks until the WaitGroup counter is zero.
	WaitGroup = sync.WaitGroup

	// 读写锁
	// Lock() 写锁，阻塞等待其它锁释放
	// RLock() 读锁
	// RUnlock() 读解锁
	// Unlock() 写解锁
	RWMutex = sync.RWMutex
)
