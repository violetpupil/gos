package slices

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestContains(t *testing.T) {
	r := slices.Contains([]string{""}, "")
	fmt.Println(r)
}

func TestReverse(t *testing.T) {
	r := []int{2, 1, 3}
	// 反转切片
	slices.Reverse[[]int, int](r)
	fmt.Println(r)
}

func TestSort(t *testing.T) {
	r := []int{2, 1, 3}
	// 升序
	slices.Sort[[]int, int](r)
	fmt.Println(r)
}

func TestSortStableFunc(t *testing.T) {
	r := []int{2, 1, 3}
	// 降序
	slices.SortStableFunc[[]int, int](r, func(a, b int) int {
		if a == b {
			return 0
		} else if a > b {
			return -1
		} else {
			return 1
		}
	})
	fmt.Println(r)
	// 升序
	slices.SortStableFunc[[]int, int](r, func(a, b int) int {
		if a == b {
			return 0
		} else if a > b {
			return 1
		} else {
			return -1
		}
	})
	fmt.Println(r)
}

func TestIndex(t *testing.T) {
	// 返回第一个索引，没有返回-1
	fmt.Println(slices.Index([]int{0, 1}, 1))
}
