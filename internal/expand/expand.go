package expand

import (
	"fmt"
	"github.com/jeremie-izzo/dctl/pkg/config"
	"github.com/jeremie-izzo/dctl/pkg/presets"
	"github.com/jeremie-izzo/dctl/pkg/runner"
	"strings"
)

type Result struct {
	Services []runner.Service
	Outputs  map[string]map[string]any // alias -> key -> value
}

//
//func PresetToServices(f *config.File) (Result, error) {
//
//	res := Result{Outputs: map[string]map[string]any{}}
//	for _, u := range f.Use {
//		p, ok := reg.Get(u.Name)
//		if !ok {
//			return res, fmt.Errorf("unknown preset: %s", u.Name)
//		}
//		params := p.NewParams()
//		if err := presets.BindParams(params, u.With); err != nil {
//			return res, fmt.Errorf("%s param bind: %w", u.Name, err)
//		}
//		alias := u.As
//		if alias == "" {
//			alias = p.DefaultAlias()
//		}
//		svcs, outs, err := p.Expand(alias, params)
//		if err != nil {
//			return res, err
//		}
//		res.Services = append(res.Services, svcs...)
//		res.Outputs[alias] = outs
//	}
//	return res, nil
//}

// Env templating: replace ${alias.key}
func ReplaceOutputs(s string, outs map[string]map[string]any) string {
	for alias, kv := range outs {
		for k, v := range kv {
			needle := fmt.Sprintf("${%s.%s}", alias, k)
			if strings.Contains(s, needle) {
				s = strings.ReplaceAll(s, needle, fmt.Sprintf("%v", v))
			}
		}
	}
	return s
}
