package imap

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	err := Login("outlook.office365.com:993", "", "")
	fmt.Println(err)
}

func TestLoginSocks5(t *testing.T) {
	err := LoginSocks5("outlook.office365.com:993", "", "", "localhost:1080", "", "")
	fmt.Println(err)
}
