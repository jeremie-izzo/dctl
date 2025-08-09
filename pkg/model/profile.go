package model

type Profile struct {
	Services []ServiceRef `yaml:"services"`
}

type ServiceRef struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

func (p *Profile) IsValid() bool {
	if len(p.Services) == 0 {
		return false
	}
	for _, service := range p.Services {
		if service.Name == "" || service.Path == "" {
			return false
		}
	}
	return true
}

func (p *Profile) GetServicePath(name string) string {
	for _, application := range p.Services {
		if application.Name == name {
			return application.Path
		}
	}
	return ""
}
