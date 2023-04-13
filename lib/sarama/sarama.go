package sarama

import (
	"crypto/sha256"
	"crypto/sha512"

	"github.com/Shopify/sarama"
)

const (
	// SCRAM Salted Challenge Response Authentication Mechanism
	// SASLTypeSCRAMSha256 represents the SCRAM-SHA-256 mechanism.
	SASLTypeSCRAMSha256 = "SCRAM-SHA-256"
	// SASLTypeSCRAMSha512 represents the SCRAM-SHA-512 mechanism.
	SASLTypeSCRAMSha512 = "SCRAM-SHA-512"
)

// SASLConfig 简单认证配置
// Simple Authentication and Security Layer
func SASLConfig(user, pass string) *sarama.Config {
	c := sarama.NewConfig()
	c.Net.SASL.Enable = true
	c.Net.SASL.User = user
	c.Net.SASL.Password = pass
	SCRAMSha256(c)
	return c
}

// SCRAMSha256 SCRAM-SHA-256 认证机制
func SCRAMSha256(c *sarama.Config) {
	c.Net.SASL.Mechanism = SASLTypeSCRAMSha256
	c.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
		return &SCRAMClient{HashGeneratorFcn: sha256.New}
	}
}

// SCRAMSha512 SCRAM-SHA-512 认证机制
func SCRAMSha512(c *sarama.Config) {
	c.Net.SASL.Mechanism = SASLTypeSCRAMSha512
	c.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
		return &SCRAMClient{HashGeneratorFcn: sha512.New}
	}
}
