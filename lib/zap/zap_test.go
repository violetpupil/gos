package zap

import (
	"testing"

	"go.uber.org/zap"
)

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
