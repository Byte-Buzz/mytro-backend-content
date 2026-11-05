package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

const envPrefix = "MT_CONTENT_"

// getEnv returns the value of the environment variable with the given key.
// If the variable does not exist, or its value is empty, the defaultValue is returned.
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(envPrefix + key)
	if !exists || value == "" {
		return defaultValue
	}
	return value
}

func getEnvWithoutPrefix(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return defaultValue
	}
	return value
}

// getIntEnv returns the value of the environment variable with the given key,
// parsed as an integer. If the variable does not exist, or its value
// is empty, the defaultValue is returned. If the value cannot be parsed
// as an integer, the defaultValue is returned.
func getIntEnv(key string, defaultValue int) int {
	valueStr, exists := os.LookupEnv(envPrefix + key)
	if !exists || valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// getDurationEnv returns the value of the environment variable with the given key,
// parsed as a time.Duration. If the variable does not exist, or its value
// is empty, the defaultValue is returned. If the value cannot be parsed
// as a time.Duration, the defaultValue is returned.
func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	valueStr, exists := os.LookupEnv(envPrefix + key)
	if !exists || valueStr == "" {
		return defaultValue
	}
	duration, err := time.ParseDuration(valueStr)
	if err != nil {
		return defaultValue
	}
	return duration
}

// getArrayEnv returns the value of the environment variable with the given key,
// split by commas into a slice of strings. If the variable does not exist,
// or its value is empty, the defaultValue is returned.
func getArrayEnv(key string, defaultValue []string) []string {
	valueStr, exists := os.LookupEnv(envPrefix + key)
	if !exists || valueStr == "" {
		return defaultValue
	}
	var values []string
	for v := range strings.SplitSeq(valueStr, ",") {
		if v == "" {
			continue
		}

		v := strings.TrimLeft(v, " ")

		values = append(values, v)
	}
	return values
}
