package presets

import (
	"github.com/jeremie-izzo/dctl/pkg/registry"
	"github.com/jeremie-izzo/dctl/pkg/runner"
	"strconv"
)

type PubSubParams struct {
	ProjectID string `mapstructure:"gcp_project_id"`
	Port      int    `mapstructure:"pubsub_port"`
}

type PubSub struct{}

func (PubSub) Name() string         { return "contrib.blueprint.pubsub" }
func (PubSub) DefaultAlias() string { return "pubsub" }
func (PubSub) NewParams() any {
	return &PubSubParams{
		ProjectID: "local-project",
		Port:      8085,
	}
}

func (PubSub) Expand(alias string, p any) ([]runner.Service, map[string]any, error) {
	pp := p.(*PubSubParams)
	if alias == "" {
		alias = "pubsub"
	}

	svc := runner.Service{
		Name:  alias + "_service",
		Image: "gcr.io/google.com/cloudsdktool/google-cloud-cli:emulators",
		Command: []string{
			"gcloud", "beta", "emulators", "pubsub", "start",
			"--host-port=0.0.0.0:8085",
			"--project=" + pp.ProjectID,
		},
		Ports: []runner.Port{
			{Host: pp.Port, Container: 8085},
		},
		Healthcheck: &runner.HealthSpec{
			Cmd: []string{"sh", "-c", "curl -s http://localhost:8085/v1/projects/" + pp.ProjectID + "/topics || exit 1"},
		},
	}

	outs := map[string]any{
		"pubsub_host":    "localhost",
		"pubsub_port":    strconv.Itoa(pp.Port),
		"gcp_project_id": pp.ProjectID,
		"service_name":   svc.Name,
	}

	return []runner.Service{svc}, outs, nil
}

func init() { registry.Global.RegisterPreset(PubSub{}) }
