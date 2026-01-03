package config

import (
	"errors"
	"os"
	"strings"
	"time"
)

type KeysConfig struct {
	PublicTokenLifetime  time.Duration
	PrivateTokenLifetime time.Duration

	privateKey string
	publicKey  string
}

func (k *KeysConfig) PrivateKey() string {
	key := k.privateKey
	k.privateKey = ""
	return key
}

func (k *KeysConfig) PublicKey() string {
	key := k.publicKey
	k.publicKey = ""
	return key
}
func (c *KeysConfig) LoadFromEnv() {
	c.privateKey = getEnv("KEYS_PRIVATE_KEY_PATH", "")
	c.publicKey = getEnv("KEYS_PUBLIC_KEY_PATH", "")
	c.PrivateTokenLifetime = getDurationEnv("TOKENS_REFRESH_TOKEN_LIFETIME", 30*24*time.Hour)
	c.PublicTokenLifetime = getDurationEnv("TOKENS_ACCESS_TOKEN_LIFETIME", 15*time.Minute)
}

func (c *KeysConfig) Validate() error {
	c.privateKey = strings.TrimSpace(c.privateKey)
	c.publicKey = strings.TrimSpace(c.publicKey)

	if c.privateKey == "" {
		return errors.New("keys.private_key is required")
	}
	if c.publicKey == "" {
		return errors.New("keys.public_key is required")
	}

	pub, err := os.ReadFile(c.publicKey)
	if err != nil {
		return errors.New("failed to read public key file: " + err.Error())
	}
	c.publicKey = string(pub)

	priv, err := os.ReadFile(c.privateKey)
	if err != nil {
		return errors.New("failed to read private key file: " + err.Error())
	}
	c.privateKey = string(priv)

	return nil
}
