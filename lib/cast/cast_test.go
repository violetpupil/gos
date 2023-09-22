package cast

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
)

func TestToIntSlice(t *testing.T) {
	fmt.Println(cast.ToIntSlice([]string{"1"}))
}
