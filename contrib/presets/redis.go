package presets

import (
	"github.com/jeremie-izzo/dctl/pkg/registry"
	"github.com/jeremie-izzo/dctl/pkg/runner"
	"path/filepath"
	"strconv"
)

type RedisParams struct {
	Port       int    `mapstructure:"redis_port"`
	VolumePath string `mapstructure:"volume_path"`
}

type Redis struct{}

func (Redis) Name() string         { return "contrib.blueprint.redis" }
func (Redis) DefaultAlias() string { return "redis" }
func (Redis) NewParams() any {
	return &RedisParams{
		Port:       6379,
		VolumePath: ".devdata",
	}
}

func (Redis) Expand(alias string, p any) ([]runner.Service, map[string]any, error) {
	pp := p.(*RedisParams)
	if alias == "" {
		alias = "redis"
	}

	svc := runner.Service{
		Name:  alias + "_db",
		Image: "redis:7",
		Ports: []runner.Port{
			{Host: pp.Port, Container: 6379},
		},
		Volumes: []runner.Volume{
			{HostPath: filepath.Join(pp.VolumePath, "redis"), ContainerPath: "/data"},
		},
		Healthcheck: &runner.HealthSpec{
			Cmd: []string{"redis-cli", "ping"},
		},
	}

	outs := map[string]any{
		"redis_host":   "localhost",
		"redis_port":   strconv.Itoa(pp.Port),
		"volume_path":  pp.VolumePath,
		"service_name": svc.Name,
	}

	return []runner.Service{svc}, outs, nil
}

func init() { registry.Global.RegisterPreset(Redis{}) }
