package time

import (
	"fmt"
	"testing"
	"time"
)

func TestIn(t *testing.T) {
	fmt.Println(In(time.Now(), 0))
	fmt.Println(In(time.Now(), 8))
}
