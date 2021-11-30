package github

import (
	"bytes"
	"context"

	"github.com/cardil/deviate/pkg/errors"
	"github.com/cardil/deviate/pkg/files"
	"github.com/cardil/deviate/pkg/metadata"
	"github.com/cli/cli/v2/pkg/cmd/factory"
	ghroot "github.com/cli/cli/v2/pkg/cmd/root"
)

// ErrClientFailed when client operations has failed.
var ErrClientFailed = errors.New("client failed")

// NewClient creates new client.
func NewClient(args ...string) Client {
	return Client{Args: args}
}

// Client a client for Github CLI.
type Client struct {
	Args         []string
	DisableColor bool
	ProjectDir   string
}

// Execute a Github client CLI command.
func (c Client) Execute(ctx context.Context) (*bytes.Buffer, error) {
	buildVersion := metadata.Version
	cmdFactory := factory.New(buildVersion)
	cmd := ghroot.NewCmdRoot(cmdFactory, buildVersion, "-")
	cmd.SetArgs(c.Args)
	var buff bytes.Buffer
	cmdFactory.IOStreams.Out = &buff
	cmdFactory.IOStreams.ErrOut = &buff
	if c.DisableColor {
		cmdFactory.IOStreams.SetColorEnabled(false)
	}
	runner := func() error {
		return errors.Wrap(cmd.ExecuteContext(ctx), ErrClientFailed)
	}
	if c.ProjectDir != "" {
		previousRunner := runner
		runner = func() error {
			return errors.Wrap(files.WithinDirectory(c.ProjectDir, previousRunner),
				ErrClientFailed)
		}
	}
	err := runner()
	return &buff, errors.Wrap(err, ErrClientFailed)
}
