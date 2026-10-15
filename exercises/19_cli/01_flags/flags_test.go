package cliflags

import (
	"testing"
)

func TestDefaults(t *testing.T) {
	cfg, err := ParseConfig([]string{})
	if err != nil {
		t.Fatalf("got error %v, want nil", err)
	}
	if cfg.Port != 8080 {
		t.Fatalf("got Port %v, want %v (hint: register -port with default 8080)", cfg.Port, 8080)
	}
	if cfg.Debug != false {
		t.Fatalf("got Debug %v, want %v (hint: register -debug with default false)", cfg.Debug, false)
	}
	if cfg.DataFile != "data.json" {
		t.Fatalf("got DataFile %v, want %v (hint: register -data with default \"data.json\")", cfg.DataFile, "data.json")
	}
}

func TestPortFlag(t *testing.T) {
	cfg, err := ParseConfig([]string{"-port", "9090"})
	if err != nil {
		t.Fatalf("got error %v, want nil", err)
	}
	if cfg.Port != 9090 {
		t.Fatalf("got Port %v, want %v (hint: call fs.Parse(args) so flags override defaults)", cfg.Port, 9090)
	}
	if cfg.Debug != false {
		t.Fatalf("got Debug %v, want %v", cfg.Debug, false)
	}
	if cfg.DataFile != "data.json" {
		t.Fatalf("got DataFile %v, want %v", cfg.DataFile, "data.json")
	}
}

func TestDebugFlag(t *testing.T) {
	cfg, err := ParseConfig([]string{"-debug"})
	if err != nil {
		t.Fatalf("got error %v, want nil", err)
	}
	if cfg.Debug != true {
		t.Fatalf("got Debug %v, want %v (hint: register -debug as a bool flag; a bare -debug sets it to true)", cfg.Debug, true)
	}
	if cfg.Port != 8080 {
		t.Fatalf("got Port %v, want %v", cfg.Port, 8080)
	}
	if cfg.DataFile != "data.json" {
		t.Fatalf("got DataFile %v, want %v", cfg.DataFile, "data.json")
	}
}

func TestDataFlag(t *testing.T) {
	cfg, err := ParseConfig([]string{"-data", "prod.json"})
	if err != nil {
		t.Fatalf("got error %v, want nil", err)
	}
	if cfg.DataFile != "prod.json" {
		t.Fatalf("got DataFile %v, want %v (hint: register -data as a string flag and return *data in Config.DataFile)", cfg.DataFile, "prod.json")
	}
	if cfg.Port != 8080 {
		t.Fatalf("got Port %v, want %v", cfg.Port, 8080)
	}
	if cfg.Debug != false {
		t.Fatalf("got Debug %v, want %v", cfg.Debug, false)
	}
}

func TestMultipleFlags(t *testing.T) {
	cfg, err := ParseConfig([]string{"-port", "3000", "-debug", "-data", "prod.json"})
	if err != nil {
		t.Fatalf("got error %v, want nil", err)
	}
	if cfg.Port != 3000 {
		t.Fatalf("got Port %v, want %v (hint: all three flags must be registered on the same FlagSet)", cfg.Port, 3000)
	}
	if cfg.Debug != true {
		t.Fatalf("got Debug %v, want %v", cfg.Debug, true)
	}
	if cfg.DataFile != "prod.json" {
		t.Fatalf("got DataFile %v, want %v", cfg.DataFile, "prod.json")
	}
}

func TestUnknownFlag(t *testing.T) {
	_, err := ParseConfig([]string{"-unknown"})
	if err == nil {
		t.Fatalf("got nil error, want non-nil (hint: use flag.ContinueOnError so Parse returns an error instead of calling os.Exit)")
	}
}
