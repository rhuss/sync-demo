package github

import (
	"bytes"
	"context"

	"github.com/cardil/deviate/pkg/errors"
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
}

// Execute a Github client CLI command.
func (c Client) Execute(ctx context.Context) (*bytes.Buffer, error) {
	buildVersion := metadata.Version
	cmdFactory := factory.New(buildVersion)
	cmd := ghroot.NewCmdRoot(cmdFactory, buildVersion, "-")
	cmd.SetArgs(c.Args)
	var buff bytes.Buffer
	cmdFactory.IOStreams.Out = &buff
	if c.DisableColor {
		cmdFactory.IOStreams.SetColorEnabled(false)
	}
	err := cmd.ExecuteContext(ctx)
	return &buff, errors.Wrap(err, ErrClientFailed)
}
