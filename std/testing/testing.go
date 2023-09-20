package testing

import "testing"

type (
	// 测试函数
	// 签名 TestFunc(t *testing.T)
	// 方法 TestSct_Func(t *testing.T)
	// 内部 Test_func(t *testing.T)
	T = testing.T
	// 测试入口函数
	// 同一个包内，执行任意测试函数，都会先执行入口函数
	// 签名 TestMain(m *testing.M)
	// 调用m.Run()执行测试函数
	M = testing.M
)
