package github

import (
	"fmt"
	"testing"
)

func TestSearchRepositories(t *testing.T) {
	err := SearchRepositories("go")
	fmt.Println(err)
}
