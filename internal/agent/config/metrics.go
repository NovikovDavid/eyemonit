package config

import (
	pkg "eyemonit/pkg/agent/config/metrics"
	"os"

	"gopkg.in/yaml.v3"
)

func NewMetricsConfig() (*pkg.MetricsConfig, error) {
	cfg := &pkg.MetricsConfig{}

	file, err := os.ReadFile("/etc/eyemonit/agent/metrics.yml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
