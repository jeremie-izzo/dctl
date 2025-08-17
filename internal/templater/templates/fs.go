package templates

import (
	"embed"
	"fmt"
)

//go:embed docker-compose.tmpl
var FS embed.FS

func Resolve(backend string) ([]byte, error) {
	embedPath := fmt.Sprintf("%s.tmpl", backend)
	b, err := FS.ReadFile(embedPath)
	if err != nil {
		return nil, fmt.Errorf("embedded template missing: %s: %w", embedPath, err)
	}
	return b, nil
}
