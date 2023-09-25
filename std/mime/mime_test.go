package mime

import (
	"fmt"
	"mime"
	"testing"
)

func TestExtensionsByType(t *testing.T) {
	r, err := mime.ExtensionsByType(JPEG)
	fmt.Println(r, err)
}
