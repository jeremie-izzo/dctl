package util

import (
	"github.com/mitchellh/mapstructure"
)

func ParseMap[T any](spec map[string]any) (*T, error) {
	var res T
	err := mapstructure.Decode(spec, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
