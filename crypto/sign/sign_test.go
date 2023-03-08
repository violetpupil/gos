package sign

import (
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {
	fmt.Println(Sign("595f23df", "1512041814", "d9f4aa7ea6d94faca62cd88a28fd5234"))
}
