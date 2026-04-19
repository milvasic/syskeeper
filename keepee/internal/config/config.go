// Package config handles loading and validation of keepee configuration.
package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config holds all keepee runtime configuration.
type Config struct {
	// ServerURL is the base URL of the keeper server.
	ServerURL string `yaml:"server_url"`

	// APIKey is the bearer token used to authenticate with the keeper server.
	APIKey string `yaml:"api_key"`

	// AgentID is the unique identifier assigned to this agent after registration.
	AgentID string `yaml:"agent_id"`

	// PushInterval is how often system snapshots are pushed to the keeper.
	PushInterval time.Duration `yaml:"push_interval"`

	// PingInterval is how often the agent pings the keeper for instructions.
	PingInterval time.Duration `yaml:"ping_interval"`
}

// DefaultConfig returns a Config populated with sensible defaults.
func DefaultConfig() Config {
	return Config{
		ServerURL:    "http://localhost:8080",
		PushInterval: 60 * time.Second,
		PingInterval: 30 * time.Second,
	}
}

// LoadFile reads a YAML config file and merges it over the defaults.
func LoadFile(path string) (Config, error) {
	cfg := DefaultConfig()

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("reading config file %q: %w", path, err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("parsing config file %q: %w", path, err)
	}

	return cfg, nil
}

// Validate checks that required fields are present.
func (c Config) Validate() error {
	if c.ServerURL == "" {
		return fmt.Errorf("server_url is required")
	}

	return nil
}
