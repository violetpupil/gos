package exp

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestContains(t *testing.T) {
	r := slices.Contains([]string{""}, "")
	fmt.Println(r)
}

func TestSort(t *testing.T) {
	r := []int{2, 1, 3}
	// 升序
	slices.Sort[int](r)
	fmt.Println(r)
}

func TestSortStableFunc(t *testing.T) {
	r := []int{2, 1, 3}
	// 降序
	slices.SortStableFunc[int](r, func(a, b int) bool { return a > b })
	fmt.Println(r)
	// 升序
	slices.SortStableFunc[int](r, func(a, b int) bool { return a < b })
	fmt.Println(r)
}
