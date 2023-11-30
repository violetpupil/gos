package atomic

import "sync/atomic"

type (
	// An Int32 is an atomic int32.
	// Add() atomically adds delta to x and returns the new value.
	// CompareAndSwap() 执行 compare-and-swap 操作
	// Load() 获取值
	// Store() 存储值
	// Swap() atomically stores new into x and returns the previous value.
	Int32 = atomic.Int32
)

var (
	// 执行 compare-and-swap 操作
	CompareAndSwapUint32 = atomic.CompareAndSwapUint32
)
