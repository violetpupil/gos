package logger

import (
	"context"
	"testing"
)

func TestOneLine(t *testing.T) {
	logger := OneLine()
	logger.Info(context.Background(), "hi")
}
