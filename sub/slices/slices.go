package slices

import "golang.org/x/exp/slices"

// Unique 去重
func Unique[T comparable](src []T) []T {
	dst := make([]T, 0)
	for _, s := range src {
		if !slices.Contains[[]T, T](dst, s) {
			dst = append(dst, s)
		}
	}
	return dst
}
