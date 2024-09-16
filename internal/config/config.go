package config

import (
	pkg "eyemonit_agent/pkg/config/agent"
	"os"

	"gopkg.in/yaml.v3"
)

func NewConfig() (*pkg.AgentConfig, error) {
	cfg := &pkg.AgentConfig{}

	file, err := os.ReadFile("/etc/eyemonit/agent/config.yml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
