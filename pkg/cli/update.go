package cli

import (
	"errors"

	"github.com/cardil/deviate/pkg/config"
	pkgerrors "github.com/cardil/deviate/pkg/errors"
	"github.com/cardil/deviate/pkg/git"
	"github.com/cardil/deviate/pkg/log"
	"github.com/cardil/deviate/pkg/log/color"
	"github.com/cardil/deviate/pkg/state"
	pkgupdate "github.com/cardil/deviate/pkg/update"
)

// ErrConfigurationIsInvalid when configuration is invalid.
var ErrConfigurationIsInvalid = errors.New("configuration is invalid")

// Upgrade will perform upgrade operation.
func Upgrade(logger log.Logger, projectFactory func() config.Project) error {
	st := state.New(log.LabeledLogger{
		Label:  color.Green("[deviate:update]"),
		Logger: logger,
	})
	defer st.Close()
	project, err := git.NewProject(projectFactory(), st)
	if err != nil {
		return pkgerrors.Wrap(err, ErrConfigurationIsInvalid)
	}
	cfg, err := config.New(project.Project, st, project.Repository())
	if err != nil {
		return pkgerrors.Wrap(err, ErrConfigurationIsInvalid)
	}
	st.Project = &project.Project
	st.Repository = project.Repository()
	st.Config = &cfg
	op := pkgupdate.Operation{State: st}
	return pkgerrors.Wrap(op.Run(), pkgupdate.ErrUpdateFailed)
}
