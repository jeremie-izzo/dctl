package compose

import (
	"bytes"
	"fmt"
	"github.com/jeremie-izzo/dctl/internal/templater/templates"
	"github.com/jeremie-izzo/dctl/pkg/runner"
	"strings"
	"text/template"
)

type Template struct{}

func (t *Template) Kind() string { return "docker-compose" }

func (t *Template) Render(p runner.Plan) ([]byte, error) {
	raw, err := templates.Resolve("docker-compose")
	if err != nil {
		return nil, err
	}

	tpl, err := template.New("compose").Funcs(funcMap()).Parse(string(raw))
	if err != nil {
		return nil, fmt.Errorf("parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, p); err != nil {
		return nil, fmt.Errorf("execute template: %w", err)
	}

	return buf.Bytes(), nil
}

// helpers for template
func funcMap() template.FuncMap {
	return template.FuncMap{
		"quote": func(s string) string {
			// basic YAML-safe double quotes, escape inner quotes
			return fmt.Sprintf(`"%s"`, strings.ReplaceAll(s, `"`, `\"`))
		},
	}
}
