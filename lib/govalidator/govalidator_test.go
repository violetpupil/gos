package govalidator

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
)

func TestIsURL(t *testing.T) {
	fmt.Println(govalidator.IsURL("http://localhost"))
}
