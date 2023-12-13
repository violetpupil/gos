package logrus

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNew(t *testing.T) {
	l := logrus.New()
	l.Infoln("hi")
}
