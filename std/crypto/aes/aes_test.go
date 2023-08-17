package aes

import (
	"fmt"
	"testing"
)

func TestEncryptS(t *testing.T) {
	r, err := EncryptS("", "", "")
	fmt.Println(r, err)
}
