package slices

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

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
