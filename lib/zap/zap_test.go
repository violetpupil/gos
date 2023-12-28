package zap

import (
	"testing"

	"go.uber.org/zap"
)

type Tmp struct {
	A string
}

func TestAny(t *testing.T) {
	err := InitDevelopment()
	if err != nil {
		panic(err)
	}
	zap.L().Info("test", zap.Any("tmp", Tmp{A: "a"}))
}

func TestInfo(t *testing.T) {
	err := InitDevelopment()
	if err != nil {
		panic(err)
	}
	zap.L().Info("test")
}

func TestNewDevelopmentFile(t *testing.T) {
	logger, err := NewDevelopmentFile("tmp.log", true)
	if err != nil {
		panic(err)
	}
	logger.Info("hi")
}
