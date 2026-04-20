package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/milvasic/syskeeper/keepee/internal/config"
)

func TestDefaultConfig(t *testing.T) {
	cfg := config.DefaultConfig()

	if cfg.ServerURL == "" {
		t.Error("expected non-empty default ServerURL")
	}

	if cfg.PushInterval <= 0 {
		t.Error("expected positive default PushInterval")
	}

	if cfg.PingInterval <= 0 {
		t.Error("expected positive default PingInterval")
	}
}

func TestLoadFile(t *testing.T) {
	content := `
server_url: http://keeper.example.com
api_key: test-key
push_interval: 30s
ping_interval: 15s
`
	f, err := os.CreateTemp(t.TempDir(), "config-*.yml")
	if err != nil {
		t.Fatalf("creating temp file: %v", err)
	}
	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("writing temp file: %v", err)
	}
	_ = f.Close()

	cfg, err := config.LoadFile(f.Name())
	if err != nil {
		t.Fatalf("LoadFile: %v", err)
	}

	if cfg.ServerURL != "http://keeper.example.com" {
		t.Errorf("ServerURL = %q; want %q", cfg.ServerURL, "http://keeper.example.com")
	}

	if cfg.APIKey != "test-key" {
		t.Errorf("APIKey = %q; want %q", cfg.APIKey, "test-key")
	}

	if cfg.PushInterval != 30*time.Second {
		t.Errorf("PushInterval = %v; want 30s", cfg.PushInterval)
	}

	if cfg.PingInterval != 15*time.Second {
		t.Errorf("PingInterval = %v; want 15s", cfg.PingInterval)
	}
}

func TestValidate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		cfg := config.DefaultConfig()
		if err := cfg.Validate(); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("missing server_url", func(t *testing.T) {
		cfg := config.Config{}
		if err := cfg.Validate(); err == nil {
			t.Error("expected error for empty ServerURL")
		}
	})
}
