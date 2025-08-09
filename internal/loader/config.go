package loader

import (
	"github.com/jeremie-izzo/dctl/pkg/model"
)

func LoadConfig() (*model.Config, error) {
	// TODO: MISSING INITIAL CONFIG
	// cliConfig, err := cli.Load()
	// if err != nil {
	// 	return nil, err
	// }
	config := &model.Config{
		ActiveProfile: "invoicing",
		CliDirectory:  ".",
	}

	return config, nil
}
