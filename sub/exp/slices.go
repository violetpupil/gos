package exp

import "golang.org/x/exp/slices"

// Unique 去重
func Unique(src []string) []string {
	dst := make([]string, 0)
	for _, s := range src {
		if !slices.Contains[[]string, string](dst, s) {
			dst = append(dst, s)
		}
	}
	return dst
}
