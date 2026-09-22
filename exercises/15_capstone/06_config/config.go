// Package config teaches environment-driven server settings.
package config

// Config holds server settings.
type Config struct {
	Port     int
	DataFile string
}

// ConfigFromEnv reads PORT and DATA_FILE, defaulting to 8080/tickets.json.
// TODO: os.Getenv + strconv.Atoi with fallback (import os, strconv).
func ConfigFromEnv() Config {
	return Config{}
}
