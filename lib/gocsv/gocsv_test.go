package gocsv

import (
	"fmt"
	"testing"
)

type Test struct {
	Name string `csv:"name"`
	Age  int    `csv:"age"`
}

func TestUnmarshalFile(t *testing.T) {
	var data []Test
	err := UnmarshalFile("./test_data/1.csv", &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}

func TestSetCSVReader(t *testing.T) {
	SetCSVReader('^')
	var data []Test
	err := UnmarshalFile("./test_data/2.csv", &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}

func TestMarshalFile(t *testing.T) {
	data := []Test{
		{Name: "jay", Age: 18},
		{Name: "john", Age: 19},
	}
	err := MarshalFile("tmp.csv", data)
	fmt.Println(err)
}
