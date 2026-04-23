package probe

import (
	"encoding/json"
	"fmt"
	"os"
)

// Target represents a single gRPC endpoint to monitor.
type Target struct {
	Name    string `json:"name"`
	Host    string `json:"host"`
	Service string `json:"service"`
}

// Config holds the list of targets to probe.
type Config struct {
	Targets      []Target `json:"targets"`
	IntervalSecs int      `json:"interval_seconds"`
}

// DefaultConfig returns a Config with sensible defaults.
func DefaultConfig() *Config {
	return &Config{
		IntervalSecs: 5,
	}
}

// LoadConfig reads and parses a JSON config file from the given path.
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	cfg := DefaultConfig()
	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parsing config file: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return cfg, nil
}

// Validate checks that the config is semantically valid.
func (c *Config) Validate() error {
	if len(c.Targets) == 0 {
		return fmt.Errorf("at least one target must be specified")
	}
	for i, t := range c.Targets {
		if t.Host == "" {
			return fmt.Errorf("target[%d]: host must not be empty", i)
		}
		if t.Name == "" {
			c.Targets[i].Name = t.Host
		}
	}
	if c.IntervalSecs <= 0 {
		return fmt.Errorf("interval_seconds must be positive")
	}
	return nil
}
