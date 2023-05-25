package user

import (
	"fmt"
	"testing"
)

func TestResponse_String(t *testing.T) {
	res := Response{Code: ResponseCode_SUCCESS}
	fmt.Println(res.String())
}
