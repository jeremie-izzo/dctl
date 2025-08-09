package model

type ComposeFile struct {
	Services map[string]ServiceConfig `yaml:"services"`
}

type ServiceConfig struct {
	Image       string                 `yaml:"image,omitempty"`
	Command     []string               `yaml:"command,omitempty"`
	Build       BuildConfig            `yaml:"build,omitempty"`
	Environment map[string]string      `yaml:"environment,omitempty"`
	Ports       []string               `yaml:"ports,omitempty"`
	Volumes     []string               `yaml:"volumes,omitempty"`
	Healthcheck map[string]interface{} `yaml:"healthcheck,omitempty"`
}

type HealthCheck struct {
	Test        []string `yaml:"test"`
	Interval    string   `yaml:"interval,omitempty"`
	Timeout     string   `yaml:"timeout,omitempty"`
	Retries     int      `yaml:"retries,omitempty"`
	StartPeriod string   `yaml:"start_period,omitempty"`
}

type BuildConfig struct {
	Context    string            `yaml:"context,omitempty"`
	Dockerfile string            `yaml:"dockerfile,omitempty"`
	Args       map[string]string `yaml:"args,omitempty"`
}
