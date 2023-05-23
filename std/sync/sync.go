package sync

import "sync"

type (
	// Map 并发使用的map
	// CompareAndDelete() 值相等则删除
	// CompareAndSwap() 值相等则交换新值
	// Delete() 删除条目
	// Load() 返回值
	// LoadAndDelete() 删除条目并返回值
	// LoadOrStore() 键存在直接返回值，否则先存储再返回
	// Range() 遍历条目
	// Store() 存储或更新条目
	// Swap() 交换值
	Map = sync.Map
)
