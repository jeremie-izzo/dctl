package presets

import (
	"github.com/jeremie-izzo/dctl/pkg/registry"
	"github.com/jeremie-izzo/dctl/pkg/runner"
	"path/filepath"
	"strconv"
)

type MySQLParams struct {
	Database   string `mapstructure:"mysql_database"`
	User       string `mapstructure:"mysql_user"`
	Password   string `mapstructure:"mysql_password"`
	Port       int    `mapstructure:"mysql_port"`
	VolumePath string `mapstructure:"volume_path"`
}

type MySQL struct{}

func (MySQL) Name() string         { return "contrib.blueprint.mysql" }
func (MySQL) DefaultAlias() string { return "mysql" }
func (MySQL) NewParams() any {
	return &MySQLParams{
		Database:   "appdb",
		User:       "appuser",
		Password:   "password",
		Port:       3306,
		VolumePath: ".devdata",
	}
}

func (MySQL) Expand(alias string, p any) ([]runner.Service, map[string]any, error) {
	pp := p.(*MySQLParams)
	if alias == "" {
		alias = "mysql"
	}

	svc := runner.Service{
		Name:  alias + "_db",
		Image: "mysql:8.0",
		Command: []string{
			"--character-set-server=utf8mb4",
			"--collation-server=utf8mb4_unicode_ci",
			"--default-authentication-plugin=mysql_native_password",
		},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "root",
			"MYSQL_DATABASE":      pp.Database,
			"MYSQL_USER":          pp.User,
			"MYSQL_PASSWORD":      pp.Password,
		},
		Ports: []runner.Port{
			{Host: pp.Port, Container: 3306},
		},
		Volumes: []runner.Volume{
			{HostPath: filepath.Join(pp.VolumePath, "mysql"), ContainerPath: "/var/lib/mysql"},
		},
		Healthcheck: &runner.HealthSpec{
			Cmd: []string{"mysqladmin", "ping", "-h", "localhost"},
		},
	}

	outs := map[string]any{
		"mysql_host":     "localhost",
		"mysql_port":     strconv.Itoa(pp.Port),
		"mysql_database": pp.Database,
		"mysql_user":     pp.User,
		"mysql_password": pp.Password,
		"volume_path":    pp.VolumePath,
		"service_name":   svc.Name,
	}

	return []runner.Service{svc}, outs, nil
}

func init() { registry.Global.RegisterPreset(MySQL{}) }
