package config

import "strings"

type CORSConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

func (c *CORSConfig) LoadFromEnv() {
	c.AllowedOrigins = getArrayEnv("CORS_ALLOWED_ORIGINS", []string{"*"})
	c.AllowedMethods = getArrayEnv("CORS_ALLOWED_METHODS", []string{
		"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
	})
	c.AllowedHeaders = getArrayEnv("CORS_ALLOWED_HEADERS", []string{
		"Origin", "Content-Type", "Authorization", "X-Requested-With",
	})
	c.AllowCredentials = strings.ToLower(getEnv("CORS_ALLOW_CREDENTIALS", "false")) == "true"
}

func (c *CORSConfig) Validate() error {
	// Можно добавить проверки, если нужны не пустые списки и допустимые методы
	if len(c.AllowedMethods) == 0 {
		c.AllowedMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	}
	if len(c.AllowedHeaders) == 0 {
		c.AllowedHeaders = []string{"Content-Type", "Authorization", "X-Requested-With"}
	}
	if len(c.AllowedOrigins) == 0 {
		c.AllowedOrigins = []string{"*"}
	}
	return nil
}
