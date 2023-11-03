package types

import (
	"fmt"
	"testing"
)

func TestHundred(t *testing.T) {
	fmt.Println(Hundred(19500))
}

func TestHundredThousand(t *testing.T) {
	fmt.Println(HundredThousand(5300000))
}

func TestHundredMillion(t *testing.T) {
	fmt.Println(HundredMillion(5300000000))
}
