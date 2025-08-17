package plugin

import (
	"fmt"
	"github.com/jeremie-izzo/dctl/internal/templater"
	"github.com/jeremie-izzo/dctl/pkg/contracts"
	"github.com/jeremie-izzo/dctl/pkg/runner"
	"github.com/jeremie-izzo/dctl/pkg/tilt"
)

type ComposePlugin struct{}

func (p *ComposePlugin) Name() string {
	return "compose"
}

func (p *ComposePlugin) Resources(_ string, _ runner.Service) ([]tilt.Resource, error) {
	return []tilt.Resource{}, nil
}

func (p *ComposePlugin) Finalize(services []runner.Service) (extra []tilt.Resource, arts []contracts.FileArtifact, err error) {
	tmp, ok := templater.Global.Get("docker-compose")
	if !ok {
		return nil, nil, fmt.Errorf("docker-compose is not configured")
	}

	plan := runner.Plan{Services: services}
	rendered, err := tmp.Render(plan)
	artifact := contracts.FileArtifact{
		Path: "./.dctl/docker-compose.yaml",
		Data: rendered,
	}

	resource := tilt.DockerCompose{
		File: "./.dctl/docker-compose.yaml",
	}
	return []tilt.Resource{resource}, []contracts.FileArtifact{artifact}, err
}
