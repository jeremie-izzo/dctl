package presets

import (
	"github.com/jeremie-izzo/dctl/pkg/registry"
	"github.com/jeremie-izzo/dctl/pkg/runner"
	"strconv"
)

type TemporalParams struct {
	TemporalPort    int    `mapstructure:"temporal_port"`
	TemporalUIPort  int    `mapstructure:"temporal_ui_port"`
	PostgresUser    string `mapstructure:"postgres_user"`
	PostgresPwd     string `mapstructure:"postgres_pwd"`
	PostgresDB      string `mapstructure:"postgres_db"`
	PostgresVersion string `mapstructure:"postgres_version"`
	TemporalVersion string `mapstructure:"temporal_version"`
	TemporalUIVer   string `mapstructure:"temporal_ui_version"`
}

type Temporal struct{}

func (Temporal) Name() string         { return "contrib.blueprint.temporal" }
func (Temporal) DefaultAlias() string { return "temporal" }
func (Temporal) NewParams() any {
	return &TemporalParams{
		TemporalPort:    7233,
		TemporalUIPort:  8080,
		PostgresUser:    "temporal",
		PostgresPwd:     "temporal",
		PostgresDB:      "temporal",
		PostgresVersion: "17",
		TemporalVersion: "1.27.2",
		TemporalUIVer:   "v2.37.4",
	}
}

func (Temporal) Expand(alias string, p any) ([]runner.Service, map[string]any, error) {
	pp := p.(*TemporalParams)
	if alias == "" {
		alias = "temporal"
	}

	pgName := alias + "_postgresql_db"
	temporalName := alias + "_service"
	uiName := alias + "_ui"

	pg := runner.Service{
		Name:  pgName,
		Image: "postgres:" + pp.PostgresVersion,
		Env: map[string]string{
			"POSTGRES_PASSWORD": pp.PostgresPwd,
			"POSTGRES_USER":     pp.PostgresUser,
			"PGUSER":            pp.PostgresUser,
		},
		Healthcheck: &runner.HealthSpec{
			Cmd: []string{"pg_isready", "-U", pp.PostgresUser},
		},
	}

	temporal := runner.Service{
		Name:  temporalName,
		Image: "temporalio/auto-setup:" + pp.TemporalVersion,
		Env: map[string]string{
			"DB":                   "postgres12",
			"DB_PORT":              "5432",
			"POSTGRES_USER":        pp.PostgresUser,
			"POSTGRES_PWD":         pp.PostgresPwd,
			"POSTGRES_SEEDS":       "postgresql",
			"ENABLE_ES":            "true",
			"ES_SEEDS":             "elasticsearch",
			"ES_VERSION":           "v7",
			"TEMPORAL_ADDRESS":     "temporal:" + strconv.Itoa(pp.TemporalPort),
			"TEMPORAL_CLI_ADDRESS": "temporal:" + strconv.Itoa(pp.TemporalPort),
		},
		Ports: []runner.Port{{Host: pp.TemporalPort, Container: 7233}},
		Healthcheck: &runner.HealthSpec{
			Cmd: []string{"pgrep", "-f", "temporal"},
		},
		DependsOn: []string{pgName},
	}

	ui := runner.Service{
		Name:  uiName,
		Image: "temporalio/ui:" + pp.TemporalUIVer,
		Env: map[string]string{
			"TEMPORAL_ADDRESS":      "temporal:" + strconv.Itoa(pp.TemporalPort),
			"TEMPORAL_CORS_ORIGINS": "http://localhost:3000",
		},
		Ports: []runner.Port{{Host: pp.TemporalUIPort, Container: 8080}},
		Healthcheck: &runner.HealthSpec{
			Cmd: []string{"wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:8080/api/health"},
		},
		DependsOn: []string{temporalName},
	}

	outs := map[string]any{
		"temporal_port":       strconv.Itoa(pp.TemporalPort),
		"temporal_ui_port":    strconv.Itoa(pp.TemporalUIPort),
		"postgres_user":       pp.PostgresUser,
		"postgres_pwd":        pp.PostgresPwd,
		"postgres_db":         pp.PostgresDB,
		"postgres_version":    pp.PostgresVersion,
		"temporal_version":    pp.TemporalVersion,
		"temporal_ui_version": pp.TemporalUIVer,
	}

	return []runner.Service{pg, temporal, ui}, outs, nil
}

func init() { registry.Global.RegisterPreset(Temporal{}) }
