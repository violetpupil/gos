package template

import (
	"fmt"
	"os"
	"testing"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func TestNew(t *testing.T) {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	fmt.Println(err)
}
