package configs

import (
	"fmt"
	"ipocalc/ipocalc/internal/models"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig(filename string) (*models.Config, error) {
	var cfg models.Config
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("YAML error: %w", err)
	}

	// default config
	if cfg.Server.Port == 0 {
		cfg.Server.Port = 8080
	}
	if cfg.Server.Host == "" {
		cfg.Server.Host = "0.0.0.0"
	}
	return &cfg, nil
}
