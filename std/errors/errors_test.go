package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrors_Error(t *testing.T) {
	var es Errors
	es = append(es, errors.New("e"))
	es = append(es, errors.New("e"))
	fmt.Println(es.Error())
}
