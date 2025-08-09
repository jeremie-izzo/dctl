package builder

import "github.com/jeremie-izzo/dctl/internal/runner"

type DeploymentManifestBuilder[T any] interface {
	Build() (T, error)
}

func NewDeploymentManifestBuilder[T any](ctx runner.RunContext) DeploymentManifestBuilder[*T] {
	return ComposeBuilder{ctx: ctx}
}

type Parser[T any] interface {
	ParseMap(spec map[string]any) (*T, error)
}
