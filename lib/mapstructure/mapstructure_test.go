package mapstructure

import (
	"fmt"
	"testing"

	"github.com/mitchellh/mapstructure"
)

type Tmp struct {
	Name string `mapstructure:"name"`
}

func TestDecode(t *testing.T) {
	m := map[string]any{"name": "jay"}
	var tmp Tmp
	err := mapstructure.Decode(m, &tmp)
	fmt.Println(tmp, err)
}
