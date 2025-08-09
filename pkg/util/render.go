package util

import (
	"strings"
	"text/template"
)

func RenderTemplate(name, tmplStr string, data any, funcs map[string]any) (string, error) {
	tmpl, err := template.New(name).Funcs(funcs).Parse(tmplStr)
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
