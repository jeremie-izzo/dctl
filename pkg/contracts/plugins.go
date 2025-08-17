package contracts

import (
	"github.com/jeremie-izzo/dctl/pkg/runner"
	"github.com/jeremie-izzo/dctl/pkg/tilt"
)

type Plugin interface {
	Name() string
	Resources(svcName string, svc runner.Service) ([]tilt.Resource, error)
}

type Aggregator interface {
	Finalize(workspace string, services map[string]runner.Service) (extra []tilt.Resource, arts []FileArtifact, err error)
}
