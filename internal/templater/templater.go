package templater

import (
	"github.com/jeremie-izzo/dctl/pkg/runner"
)

type Kind string

type Templater interface {
	Kind() string
	Render(p runner.Plan) ([]byte, error)
}
