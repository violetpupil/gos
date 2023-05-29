package govalidator

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
)

func TestIsDNSName(t *testing.T) {
	fmt.Println(govalidator.IsDNSName("www.baidu.com"))
}

func TestIsURL(t *testing.T) {
	fmt.Println(govalidator.IsURL("http://localhost"))
}
