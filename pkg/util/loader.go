package util

import (
	"gopkg.in/yaml.v3"
	"os"
)

func LoadYaml[T any](path string) (*T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var result *T
	if err = yaml.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
