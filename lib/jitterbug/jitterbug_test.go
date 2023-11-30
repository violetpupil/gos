package jitterbug

import (
	"fmt"
	"testing"
	"time"

	"github.com/lthibault/jitterbug/v2"
)

func TestNorm(t *testing.T) {
	n := jitterbug.Norm{Stdev: 1 * time.Second}
	fmt.Println(n.Jitter(5 * time.Second))
}
