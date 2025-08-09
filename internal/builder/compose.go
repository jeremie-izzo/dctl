package builder

import (
	"github.com/jeremie-izzo/dctl/internal/resolver"
	"github.com/jeremie-izzo/dctl/internal/runner"
	"github.com/jeremie-izzo/dctl/pkg/model"
	"github.com/jeremie-izzo/dctl/pkg/util"
)

type ComposeBuilder struct {
	ctx runner.RunContext
}

func (c ComposeBuilder) Build() (*model.ComposeFile, error) {
	final := &model.ComposeFile{
		Services: map[string]model.ServiceConfig{},
	}

	for _, instruction := range c.ctx.Instructions {
		svcCfg, err := util.ParseMap[model.ServiceConfig](instruction.Deployment.RawConfig)
		if err != nil {
			return nil, err
		}
		final.Services[instruction.Name] = *svcCfg

		for _, dep := range instruction.Dependencies {
			depPath, err := resolver.DependencyFilePath(dep.Name, dep.Version, instruction.Type, c.ctx.Config.CliDirectory)
			if err != nil {
				return nil, err
			}

			depCfg, err := util.LoadYaml[model.ComposeFile](depPath)
			if err != nil {
				return nil, err
			}
			for name, svc := range depCfg.Services {
				final.Services[name] = svc
			}
		}
	}

	return final, nil
}
