package pkg

type Metrics struct {
	Cpu     bool `yaml:"cpu"`
	Memory  bool `yaml:"memory"`
	Disk    bool `yaml:"disk"`
	Network bool `yaml:"network"`
}

type MetricsConfig struct {
	Metrics Metrics `yaml:"metrics"`
}
