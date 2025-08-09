package orchestrator

import "github.com/jeremie-izzo/dctl/internal/runner"

func RunTemplate() (*str, error) {
	runContext, err := runner.InitContext()
	if err != nil {
		return nil, err
	}

	composeFile, err := builder.ComposeFileFromProfile(runContext)
	if err != nil {
		return nil, err
	}

	return "string", nil
}
