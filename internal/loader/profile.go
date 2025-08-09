package loader

import (
	"fmt"
	"github.com/jeremie-izzo/dctl/internal/resolver"
	"github.com/jeremie-izzo/dctl/pkg/model"
	"github.com/jeremie-izzo/dctl/pkg/util"
)

func LoadProfile(conf *model.Config) (*model.Profile, error) {
	profilePath, err := resolver.ProfileFilePath(conf.CliDirectory, conf.ActiveProfile)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve profile path: %w", err)
	}

	pro, err := util.LoadYaml[model.Profile](profilePath)
	if err != nil {
		return nil, fmt.Errorf("unable to load profile %s: %w", profilePath, err)
	}

	if !pro.IsValid() {
		return nil, fmt.Errorf("profile %s is not valid", conf.ActiveProfile)
	}

	return pro, nil
}
