package gofakeit

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestName(t *testing.T) {
	fmt.Println(gofakeit.Name())
}

func TestIntRange(t *testing.T) {
	fmt.Println(gofakeit.IntRange(0, 10))
}

func TestUserAgent(t *testing.T) {
	fmt.Println(gofakeit.UserAgent())
}
