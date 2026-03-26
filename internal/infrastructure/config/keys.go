package config

import (
	"errors"
	"os"
	"strings"
)

type KeysConfig struct {
	publicKey string
}

func (k *KeysConfig) PublicKey() string {
	key := k.publicKey
	k.publicKey = ""
	return key
}
func (c *KeysConfig) LoadFromEnv() {
	c.publicKey = getEnv("KEYS_PUBLIC_KEY_PATH", "")
}

func (c *KeysConfig) Validate() error {
	c.publicKey = strings.TrimSpace(c.publicKey)

	if c.publicKey == "" {
		return errors.New("keys.public_key is required")
	}

	pub, err := os.ReadFile(c.publicKey)
	if err != nil {
		return errors.New("failed to read public key file: " + err.Error())
	}
	c.publicKey = string(pub)

	return nil
}
