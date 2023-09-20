package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStdout(t *testing.T) {
	Stdout(map[string]string{"a": "b"})
}

type Tmp struct {
	A bool `json:",string"`
}

func TestMarshal(t *testing.T) {
	var tmp Tmp
	bs, err := json.Marshal(tmp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}
