package resolver

import (
	"fmt"
	"github.com/jeremie-izzo/dctl/pkg/model"
	"path/filepath"
)

func ProfileFilePath(basePath, profileName string) (string, error) {
	if profileName == "" {
		return "", fmt.Errorf("profile name cannot be empty")
	}
	return filepath.Join(basePath, "profiles", profileName+".yaml"), nil
}

func DependencyFilePath(name, version string, runtime model.Runtime, basePath string) (string, error) {
	switch runtime {
	case model.TypeCompose:
		return filepath.Join(basePath, "dependencies", name, version, "compose", "docker-compose.yaml"), nil
	case model.TypeK8s:
		return filepath.Join(basePath, "dependencies", name, version, "k8s", "deployment.yaml"), nil
	default:
		return "", fmt.Errorf("unsupported runtime: %s", runtime)
	}
}

func InstructionFilePath(basePath string) string {
	return filepath.Join(basePath, "dctl.yaml")
}

func TemplateFilePath(basePath string) string {
	return filepath.Join(basePath, "templates", "docker-compose.tmpl")
}
