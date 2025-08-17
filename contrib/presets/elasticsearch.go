package presets

import (
	"github.com/jeremie-izzo/dctl/pkg/registry"
	"github.com/jeremie-izzo/dctl/pkg/runner"
	"path/filepath"
	"strconv"
)

type ESParams struct {
	// Versions
	ElasticsearchVersion string `mapstructure:"elasticsearch_version"` // e.g., "8.17.0"
	KibanaVersion        string `mapstructure:"kibana_version"`        // e.g., "8.17.0"

	// Ports (host)
	ElasticsearchPort int `mapstructure:"elasticsearch_port"` // -> 9200 in container
	TransportPort     int `mapstructure:"transport_port"`     // -> 9300 in container
	KibanaPort        int `mapstructure:"kibana_port"`        // -> 5601 in container

	// Data dir
	VolumePath string `mapstructure:"volume_path"`
}

type ES struct{}

func (ES) Name() string         { return "contrib.blueprint.elasticsearch" }
func (ES) DefaultAlias() string { return "elasticsearch" }
func (ES) NewParams() any {
	return &ESParams{
		ElasticsearchVersion: "8.17.0",
		KibanaVersion:        "8.17.0",
		ElasticsearchPort:    9200,
		TransportPort:        9300,
		KibanaPort:           5601,
		VolumePath:           ".devdata",
	}
}

func (ES) Expand(alias string, p any) ([]runner.Service, map[string]any, error) {
	pp := p.(*ESParams)
	if alias == "" {
		alias = "elasticsearch"
	}

	esName := alias + "_db"
	kbName := alias + "_ui"

	// Elasticsearch node (single-node, no security)
	es := runner.Service{
		Name:  esName,
		Image: "docker.elastic.co/elasticsearch/elasticsearch:" + pp.ElasticsearchVersion,
		Env: map[string]string{
			"discovery.type":                       "single-node",
			"ES_JAVA_OPTS":                         "-Xms512m -Xmx512m",
			"xpack.security.enabled":               "false",
			"xpack.security.enrollment.enabled":    "false",
			"xpack.security.http.ssl.enabled":      "false",
			"xpack.security.transport.ssl.enabled": "false",
		},
		Ports: []runner.Port{
			{Host: pp.ElasticsearchPort, Container: 9200},
			{Host: pp.TransportPort, Container: 9300},
		},
		Volumes: []runner.Volume{
			{HostPath: filepath.Join(pp.VolumePath, "elasticsearch"), ContainerPath: "/usr/share/elasticsearch/data"},
		},
		Healthcheck: &runner.HealthSpec{
			// Compose will emit CMD form; use sh -c for grep pipeline
			Cmd: []string{"sh", "-c", "curl -s http://localhost:9200 | grep -q cluster_name"},
		},
	}

	// Kibana UI
	kb := runner.Service{
		Name:  kbName,
		Image: "kibana:" + pp.KibanaVersion,
		Env: map[string]string{
			"ELASTICSEARCH_URL":   "http://" + esName + ":9200",
			"ELASTICSEARCH_HOSTS": "http://" + esName + ":9200",
		},
		Ports: []runner.Port{
			{Host: pp.KibanaPort, Container: 5601},
		},
		Healthcheck: &runner.HealthSpec{
			Cmd: []string{"sh", "-c", "curl -s http://localhost:5601 || exit 1"},
		},
		DependsOn: []string{esName},
	}

	outs := map[string]any{
		"es_url":       "http://localhost:" + strconv.Itoa(pp.ElasticsearchPort),
		"kibana_url":   "http://localhost:" + strconv.Itoa(pp.KibanaPort),
		"es_service":   esName,
		"kib_service":  kbName,
		"data_path":    pp.VolumePath,
		"es_version":   pp.ElasticsearchVersion,
		"kib_version":  pp.KibanaVersion,
		"transportNic": strconv.Itoa(pp.TransportPort),
	}

	return []runner.Service{es, kb}, outs, nil
}

func init() { registry.Global.RegisterPreset(ES{}) }
