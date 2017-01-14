package valid

import (
	"errors"

	"github.com/goreleaser/releaser/context"
)

// Pipe for brew deployment
type Pipe struct{}

// Name of the pipe
func (Pipe) Name() string {
	return "Valid"
}

// Run the pipe
func (Pipe) Run(ctx *context.Context) (err error) {
	if ctx.Config.BinaryName == "" {
		return errors.New("missing binary_name")
	}
	if ctx.Config.Repo == "" {
		return errors.New("missing repo")
	}
	return
}
