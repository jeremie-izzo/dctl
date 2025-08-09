package runner

import (
	"github.com/jeremie-izzo/dctl/internal/loader"
	"github.com/jeremie-izzo/dctl/pkg/model"
)

type RunContext struct {
	Config       *model.Config
	Profile      *model.Profile
	Instructions []model.InstructionManifest
}

func InitContext() (*RunContext, error) {
	cliConfig, err := loader.LoadConfig()
	if err != nil {
		return nil, err
	}

	builtProfile, err := loader.LoadProfile(cliConfig)
	if err != nil {
		return nil, err
	}

	builtInstructions, err := loader.LoadInstructions(builtProfile)
	if err != nil {
		return nil, err
	}

	return &RunContext{
		Config:       cliConfig,
		Profile:      builtProfile,
		Instructions: builtInstructions,
	}, nil
}
