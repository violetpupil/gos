package crypto

import (
	"crypto/ecdh"
	"crypto/rand"
	"encoding/base64"
	"testing"

	"github.com/violetpupil/gos/misc"
	"go.uber.org/zap"
)

func TestECDH(t *testing.T) {
	misc.InitDevelopmentLog()
	log := zap.L()

	cliKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal("generate client key error", zap.Error(err))
	}
	cliPub := base64.StdEncoding.EncodeToString(cliKey.PublicKey().Bytes())

	ECDH(log, cliPub)
}
