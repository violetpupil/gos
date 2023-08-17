package md5

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	fmt.Println(Sum("595f23df1512041814"))
}

func TestSumSalt(t *testing.T) {
	fmt.Println(SumSalt("abc", "123"))
}
