// Package config teaches environment-driven server settings.
package config

import (
	"os"
	"strconv"
)

// Config holds server settings.
type Config struct {
	Port     int
	DataFile string
}

// ConfigFromEnv reads PORT and DATA_FILE, defaulting to 8080/tickets.json.
func ConfigFromEnv() Config {
	cfg := Config{Port: 8080, DataFile: "tickets.json"}
	if v := os.Getenv("PORT"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			cfg.Port = n
		}
	}
	if v := os.Getenv("DATA_FILE"); v != "" {
		cfg.DataFile = v
	}
	return cfg
}
