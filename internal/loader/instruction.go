package loader

import (
	"fmt"
	"github.com/jeremie-izzo/dctl/internal/resolver"
	"github.com/jeremie-izzo/dctl/pkg/model"
	"github.com/jeremie-izzo/dctl/pkg/util"
)

func LoadInstructions(profile *model.Profile) ([]model.InstructionManifest, error) {
	var instructions []model.InstructionManifest

	var runtime model.Runtime
	for _, svc := range profile.Services {
		path := resolver.InstructionFilePath(svc.Path)

		manifest, err := util.LoadYaml[model.InstructionManifest](path)
		if err != nil {
			return nil, fmt.Errorf("load instruction manifest %s: %v", svc.Path, err)
		}

		instructions = append(instructions, *manifest)
	}

	return instructions, nil
}
