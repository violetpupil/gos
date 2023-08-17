package hmac

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	fmt.Printf("%x\n", Sum(sha1.New, "", ""))
	fmt.Printf("%x\n", Sum(sha256.New, "", ""))
	fmt.Printf("%x\n", Sum(sha512.New, "", ""))
}
