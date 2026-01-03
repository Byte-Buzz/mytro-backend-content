package config

import "errors"

type LoggingConfig struct {
	Level    string
	Format   string
	Output   string
	FilePath string
}

func (c *LoggingConfig) LoadFromEnv() {
	c.Level = getEnv("LOGGING_LEVEL", "info")
	c.Format = getEnv("LOGGING_FORMAT", "json")
	c.Output = getEnv("LOGGING_OUTPUT", "stdout")
	c.FilePath = getEnv("LOGGING_FILE_PATH", "")
}

func (c *LoggingConfig) Validate() error {
	switch c.Level {
	case "debug", "info", "warn", "error":
	default:
		return errors.New("logging.level must be one of debug, info, warn, error")
	}
	switch c.Format {
	case "json", "console":
	default:
		return errors.New("logging.format must be one of json, console")
	}
	switch c.Output {
	case "stdout", "stderr":
	default:
		return errors.New("logging.output must be one of stdout, stderr")
	}
	return nil
}
