package presets

import (
	"github.com/jeremie-izzo/dctl/pkg/registry"
	"github.com/jeremie-izzo/dctl/pkg/runner"
	"path/filepath"
	"strconv"
)

type DSParams struct {
	ProjectID  string `mapstructure:"project_id"`
	Port       int    `mapstructure:"port"`    // host port for emulator (-> 8081 in container)
	UIPort     int    `mapstructure:"ui_port"` // host port for UI (-> 8080 in container)
	VolumePath string `mapstructure:"volume_path"`
}

type DS struct{}

func (DS) Name() string         { return "contrib.blueprint.datastore" }
func (DS) DefaultAlias() string { return "datastore" }
func (DS) NewParams() any {
	return &DSParams{
		ProjectID:  "local-dev",
		Port:       8081,
		UIPort:     8080,
		VolumePath: ".devdata",
	}
}

func (DS) Expand(alias string, p any) ([]runner.Service, map[string]any, error) {
	pp := p.(*DSParams)
	if alias == "" {
		alias = "datastore"
	}

	// datastore emulator service (gcloud)
	db := runner.Service{
		Name:  alias, // e.g., "datastore"
		Image: "gcr.io/google.com/cloudsdktool/google-cloud-cli:522.0.0-emulators",
		Command: []string{
			"gcloud", "beta", "emulators", "datastore", "start",
			"--host-port=0.0.0.0:8081",
			"--project=" + pp.ProjectID,
			"--data-dir=/data",
		},
		Ports: []runner.Port{
			{Host: pp.Port, Container: 8081},
		},
		Volumes: []runner.Volume{
			{HostPath: filepath.Join(pp.VolumePath, "datastore"), ContainerPath: "/data", ReadOnly: false},
		},
		Env: map[string]string{},
		Healthcheck: &runner.HealthSpec{
			// compose template will turn this into ["CMD", ...]
			Cmd: []string{"curl", "-s", "http://localhost:8081"},
		},
	}

	// datastore admin UI (dsadmin)
	ui := runner.Service{
		Name:    alias + "-ui", // e.g., "datastore-ui"
		Image:   "ghcr.io/remko/dsadmin:latest",
		Command: nil,
		Ports: []runner.Port{
			{Host: pp.UIPort, Container: 8080},
		},
		Env: map[string]string{
			"DATASTORE_PROJECT_ID":    pp.ProjectID,
			"DATASTORE_EMULATOR_HOST": alias + ":8081", // reach emulator by service name
		},
	}

	outs := map[string]any{
		"emulator_host":   "localhost:",
		"ui_host":         "localhost:" + strconv.Itoa(pp.UIPort),
		"project_id":      pp.ProjectID,
		"volume_path":     pp.VolumePath,
		"service_name":    alias,
		"ui_service_name": alias + "-ui",
	}

	return []runner.Service{db, ui}, outs, nil
}

// Side-effect registration for easy inclusion via blank import
func init() { registry.Global.RegisterPreset(DS{}) }
