package cmd

import (
	"errors"
	"path"

	"github.com/cardil/deviate/pkg/cli"
	"github.com/cardil/deviate/pkg/config"
	"github.com/spf13/cobra"
)

var (
	// ErrConfigurationIsInvalid when configuration is invalid.
	ErrConfigurationIsInvalid = errors.New("configuration is invalid")
	// ErrUpdateFailed when an update fails.
	ErrUpdateFailed = errors.New("update failed")
)

type update struct {
	*cli.Options
}

func (u update) command() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "update",
		Short:     "Synchronize to the upstream releases",
		ValidArgs: []string{"REPOSITORY"},
		Args:      cobra.OnlyValidArgs,
		RunE:      u.run,
	}
	return cmd
}

func (u update) run(cmd *cobra.Command, args []string) error {
	return cli.Upgrade(cmd, u.project(args)) //nolint:wrapcheck
}

func (u update) project(args []string) func() config.Project {
	return func() config.Project {
		project := config.Project{
			ConfigPath: u.ConfigPath,
			Path:       path.Dir(u.ConfigPath),
		}
		if len(args) > 0 {
			project.Path = args[0]
		}
		return project
	}
}
