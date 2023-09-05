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

func (*Fruit) ReadConfigure(*exl.ReadConfig) {}

func (*Fruit) WriteConfigure(*exl.WriteConfig) {}

func TestReadFile(t *testing.T) {
	fruits, err := exl.ReadFile[*Fruit]("./test_data/fruits.xlsx")
	if err != nil {
		panic(err)
	}
	for _, f := range fruits {
		fmt.Printf("%+v\n", f)
	}

	err = exl.Write[*Fruit]("tmp.xlsx", fruits)
	fmt.Println(err)
}
