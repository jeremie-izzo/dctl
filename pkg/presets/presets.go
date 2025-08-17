package presets

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/jeremie-izzo/dctl/pkg/runner"
)

type ServicePresets interface {
	Name() string
	DefaultAlias() string
	NewParams() any // pointer to a params struct with defaults
	Expand(alias string, p any) (services []runner.Service, outputs map[string]any, err error)
}

// BindParams maps a free-form map into a typed params struct with defaults.
func BindParams(dst any, src map[string]any) error { return mapstructure.Decode(src, dst) }
