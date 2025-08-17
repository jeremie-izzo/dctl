package runner

import (
	"time"
)

type Plan struct {
	Services []Service
}

type Service struct {
	Name        string            `yaml:"name"`
	Image       string            `yaml:"image,omitempty"`
	Build       *BuildSpec        `yaml:"build,omitempty"`       // optional, if not provided, image is used directly
	Command     []string          `yaml:"command,omitempty"`     // optional, if not provided, defaults to the image's CMD
	Env         map[string]string `yaml:"env,omitempty"`         // optional, if not provided, defaults to the image's environment
	Ports       []Port            `yaml:"ports,omitempty"`       // optional, if not provided, no ports are exposed
	Volumes     []Volume          `yaml:"volumes,omitempty"`     // optional, if not provided, no volumes are mounted
	DependsOn   []string          `yaml:"depends_on,omitempty"`  // optional, if not provided, no dependencies
	Healthcheck *HealthSpec       `yaml:"healthcheck,omitempty"` // optional, if not provided, no healthcheck is defined
	Plugins     []string          `yaml:"plugins,omitempty"`
}

type BuildSpec struct {
	Context    string            `yaml:"context"`
	Dockerfile string            `yaml:"dockerfile,omitempty"` // optional, defaults to "Dockerfile" in the context directory
	Args       map[string]string `yaml:"args,omitempty"`       // optional, build arguments
}

type Port struct{ Host, Container int }

type Volume struct {
	HostPath, ContainerPath string
	ReadOnly                bool
}

type HealthSpec struct {
	HTTPGet   string        // optional
	TCPSocket string        // host:port optional
	Cmd       []string      // optional
	Interval  time.Duration // default 2s
	Timeout   time.Duration // default 500ms
	Retries   int           // default 30
}
