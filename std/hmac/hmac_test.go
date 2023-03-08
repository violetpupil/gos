package hmac

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	fmt.Println(Sum(sha1.New, "", ""))
	fmt.Println(Sum(sha256.New, "", ""))
	fmt.Println(Sum(sha512.New, "", ""))
}
