package cli

import (
	"errors"

	"github.com/cardil/deviate/pkg/config"
	"github.com/cardil/deviate/pkg/git"
	"github.com/cardil/deviate/pkg/log"
	"github.com/cardil/deviate/pkg/state"
	pkgupdate "github.com/cardil/deviate/pkg/update"
)

var (
	// ErrConfigurationIsInvalid when configuration is invalid.
	ErrConfigurationIsInvalid = errors.New("configuration is invalid")
	// ErrUpdateFailed when an update fails.
	ErrUpdateFailed = errors.New("update failed")
)

// Upgrade will perform upgrade operation.
func Upgrade(log log.Logger, projectFactory func() config.Project) error {
	st := state.New(log)
	defer st.Close()
	project, err := git.New(projectFactory())
	if err != nil {
		return wrap(err, ErrConfigurationIsInvalid)
	}
	cfg, err := config.New(project)
	if err != nil {
		return wrap(err, ErrConfigurationIsInvalid)
	}
	st.Repository = project.Repository()
	st.Config = &cfg
	op := pkgupdate.Operation{State: st}
	return wrap(op.Run(), ErrUpdateFailed)
}
