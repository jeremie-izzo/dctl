package model

type Runtime string

const (
	TypeCompose Runtime = "compose"
	TypeK8s     Runtime = "k8s"
)

type InstructionManifest struct {
	Name         string         `yaml:"name"`
	Type         Runtime        `yaml:"runtime"`
	Deployment   DeploymentSpec `yaml:"deployment"`
	Dependencies []Dependency   `yaml:"dependencies,omitempty"`
	Commands     []Command      `yaml:"commands,omitempty"`
}

type DeploymentSpec struct {
	RawConfig map[string]any `yaml:"configuration"`
}

type Command struct {
	Name string `yaml:"name"`
	Run  string `yaml:"run"`
}

type Dependency struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Runtime string `yaml:"runtime"`
}
