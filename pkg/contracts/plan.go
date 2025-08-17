package contracts

type FileArtifact struct {
	Path string // local path to write to (workspace)
	Data []byte // content
}

type Plan struct {
	Name      string
	Artifacts []FileArtifact // e.g., docker-compose.yml, k8s manifests
}
