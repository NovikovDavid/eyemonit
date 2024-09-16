package pkg

type Agent struct {
	Name       string `yaml:"name"`
	Interval   int    `yaml:"interval"`
	LogLevel   string `yaml:"log_level"`
	RetryCount int    `yaml:"retry_count"`
}

type Server struct {
	Port     int    `yaml:"port"`
	Address  string `yaml:"address"`
	Protocol string `yaml:"protocol"`
	Endpoint string `yaml:"endpoint"`
}

type Host struct {
	Tags  []string `yaml:"tags"`
	Group string   `yaml:"group"`
}

type Security struct {
	AuthToken string `yaml:"auth_token"`
	SSLVerify bool   `yaml:"ssl_verify"`
}

type Advanced struct {
	Timeout      int `yaml:"timeout"`
	MaxQueueSize int `yaml:"max_queue_size"`
}

type AgentConfig struct {
	ConfigAgent    Agent    `yaml:"agent"`
	ConfigServer   Server   `yaml:"server"`
	ConfigHost     Host     `yaml:"host"`
	ConfigSecurity Security `yaml:"security"`
	ConfigAdvanced Advanced `yaml:"advanced"`
}
