package config

import "testing"

func TestConfigFromEnv(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("DATA_FILE", "/tmp/t.json")
	got := ConfigFromEnv()
	if got.Port != 9090 || got.DataFile != "/tmp/t.json" {
		t.Fatalf("Config = %+v", got)
	}
}
