package contracts

type TemplateEngine interface {
	// Render transforms a DeploymentPlan into Tiltfile/Starlark and aux files.
	Render(plan Plan) (files []FileArtifact, err error)
	EngineName() string
}
