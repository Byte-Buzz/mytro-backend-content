package config

import (
	"errors"
	"time"
)

type ServerConfig struct {
	Host           string
	Port           int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	RequestTimeout time.Duration
}

func (c *ServerConfig) LoadFromEnv() {
	c.Host = getEnv("SERVER_HOST", "0.0.0.0")
	c.Port = getIntEnv("SERVER_PORT", 8080)
	c.ReadTimeout = getDurationEnv("SERVER_READ_TIMEOUT", 10*time.Second)
	c.WriteTimeout = getDurationEnv("SERVER_WRITE_TIMEOUT", 20*time.Second)
	c.RequestTimeout = getDurationEnv("SERVER_REQUEST_TIMEOUT", 10*time.Second)
}

func (c *ServerConfig) Validate() error {
	if c.Port <= 0 || c.Port > 65535 {
		return errors.New("server.port must be a number between 1 and 65535")
	}
	return nil
}
