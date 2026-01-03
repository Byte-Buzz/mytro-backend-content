package config

import (
	"errors"
	"time"
)

type DatabaseConfig struct {
	URL             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func LoadDatabaseFromEnv() (*DatabaseConfig, error) {
	cfg := &DatabaseConfig{
		URL:             getEnv("DATABASE_URL", ""),
		MaxOpenConns:    getIntEnv("DATABASE_MAX_OPEN_CONNS", 15),
		MaxIdleConns:    getIntEnv("DATABASE_MAX_IDLE_CONNS", 5),
		ConnMaxLifetime: getDurationEnv("DATABASE_CONN_MAX_LIFETIME", 1*time.Hour),
	}

	if cfg.URL == "" {
		return nil, errors.New("database.dsn is required")
	}

	return cfg, nil
}
