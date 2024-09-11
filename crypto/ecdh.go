package crypto

import (
	"crypto/ecdh"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"go.uber.org/zap"
)

// 登录
// 客户端 --请求携带客户端公钥--> 服务端，服务端使用客户端公钥和服务端私钥生成密钥
// 服务端 --响应携带服务端公钥--> 客户端，客户端使用服务端公钥和客户端私钥生成密钥
// 两端生成的密钥是相同的，各自保存密钥，之后通信用AES加密

func ECDH(log *zap.Logger, cliPub string) (svrPub string, shared string, err error) {
	svrKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		log.Error("generate key error", zap.Error(err))
		return "", "", err
	}
	cliPubO, err := ecdh.P256().NewPublicKey([]byte(cliPub))
	if err != nil {
		log.Error("new public key error", zap.Error(err))
		return "", "", err
	}
	sharedB, err := svrKey.ECDH(cliPubO)
	if err != nil {
		log.Error("generate shared error", zap.Error(err))
		return "", "", err
	}

	svrPub = fmt.Sprintf("%x", sha256.Sum256(svrKey.PublicKey().Bytes()))
	shared = fmt.Sprintf("%x", sha256.Sum256(sharedB))
	return svrPub, shared, nil
}
