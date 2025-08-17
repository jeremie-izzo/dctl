package tilt

import (
	"fmt"
	"github.com/jeremie-izzo/dctl/pkg/runner"
	"github.com/jeremie-izzo/dctl/pkg/tilt"
)

func Render(plan runner.Plan, reg tilt.Registry) error {
	b := tilt.NewEmitter()

	for _, s := range plan.Services {
		kind := "compose"
		if s.Deploy != nil && s.Deploy.Type != "" {
			kind = s.Deploy.Type
		}

		p, ok := reg.Get(kind)
		if !ok {
			return fmt.Errorf("unknown deploy type %q for service %s", kind, s.Name)
		}

		if err := p.Prepare(&s); err != nil {
			return err
		}
		if err := p.Build(s, b); err != nil {
			return err
		}
	}
	return nil
}
