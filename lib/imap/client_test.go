package imap

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	err := Login("outlook.office365.com:993", "", "")
	fmt.Println(err)
}
