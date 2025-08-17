// tilt/resources.go
package tilt

type Resource interface{ isTiltResource() }

type LocalResource struct {
	Name, Cmd    string
	Deps         []string
	TriggerMode  string   // "TRIGGER_MODE_AUTO" or "TRIGGER_MODE_MANUAL"
	ResourceDeps []string // other resources this one depends on
}

func (LocalResource) isTiltResource() {}

type DockerCompose struct {
	File    string // path to docker-compose.yml
	Project string
}

func (DockerCompose) isTiltResource() {}

type DCResource struct {
	Name         string
	TriggerMode  string   // "TRIGGER_MODE_AUTO" or "TRIGGER_MODE_MANUAL"
	ResourceDeps []string // other resources this one depends on
	Labels       map[string]string
	ProjectName  string
}

func (DCResource) isTiltResource() {}
