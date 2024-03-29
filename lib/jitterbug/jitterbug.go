package jitterbug

import (
	"time"

	"github.com/lthibault/jitterbug/v2"
)

type (
	// Norm is a normal distribution.
	// Stdev 标准差
	//
	// Jitter() 在时间上添加抖动，返回最终时间
	Norm = jitterbug.Norm
)

// NewNorm 创建正态分布抖动 ticker
func NewNorm(d time.Duration, stdev time.Duration) *jitterbug.Ticker {
	return jitterbug.New(d, &jitterbug.Norm{Stdev: stdev})
}
