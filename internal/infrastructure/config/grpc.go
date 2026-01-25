package config

import "errors"

type GRPCConfig struct {
	ContentStorageAddress string
}

func (c *GRPCConfig) LoadFromEnv() {
	c.ContentStorageAddress = getEnvWithoutPrefix("CONTENT_STORAGE_ADDR", "")
}

func (c *GRPCConfig) Validate() error {
	if c.ContentStorageAddress == "" {
		return errors.New("grpc.content_storage_addr is required")
	}
	return nil
}

func LoadGRPCFromEnv() (*GRPCConfig, error) {
	c := &GRPCConfig{}
	c.LoadFromEnv()
	return c, c.Validate()
}
