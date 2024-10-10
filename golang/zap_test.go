package golang

import (
	"testing"

	"go.uber.org/zap"
)

func TestZapAny(t *testing.T) {
	InitDevelopmentLog()
	zap.L().Info("", zap.Any("struct", struct{ I int }{}))
}
