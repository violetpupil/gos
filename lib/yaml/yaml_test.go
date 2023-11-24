package yaml

import (
	"fmt"
	"testing"
)

type Data struct {
	A string
}

func TestUnmarshalFile(t *testing.T) {
	var d Data
	err := UnmarshalFile("data.yml", &d)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", d)
}
