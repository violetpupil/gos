package jwt

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	r, err := GenToken()
	fmt.Println(r, err)
}

func TestParseToken(t *testing.T) {
	token, err := GenToken()
	if err != nil {
		panic(err)
	}
	claims, err := ParseToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+s", claims)
}
