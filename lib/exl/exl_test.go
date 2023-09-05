package exl

import (
	"fmt"
	"testing"

	"github.com/go-the-way/exl"
)

type Fruit struct {
	ID   int    `excel:"ID"`
	Name string `excel:"Name"`
}

func (*Fruit) Configure(*exl.ReadConfig) {}

func TestReadFile(t *testing.T) {
	fruits, err := exl.ReadFile[*Fruit]("./test_data/fruits.xlsx")
	if err != nil {
		panic(err)
	}
	for _, f := range fruits {
		fmt.Printf("%+v\n", f)
	}
}
