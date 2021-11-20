package sync

import (
	"github.com/cardil/deviate/pkg/config/git"
	"github.com/cardil/deviate/pkg/errors"
	"github.com/cardil/deviate/pkg/state"
)

// ErrSyncFailed when the sync failed.
var ErrSyncFailed = errors.New("sync failed")

// Operation performs sync - the upstream synchronization.
type Operation struct {
	state.State
}

func (o Operation) Run() error {
	err := runSteps([]step{
		o.mirrorReleases,
		o.syncReleaseNext,
		o.triggerCI,
		o.createPR,
	})
	_ = o.switchToMain()
	return err
}

func (o Operation) switchToMain() error {
	downstream := git.Remote{Name: "downstream", URL: o.Config.Downstream}
	err := o.Repository.Fetch(downstream)
	if err != nil {
		return errors.Wrap(err, ErrSyncFailed)
	}
	return errors.Wrap(
		o.Repository.Checkout(downstream, o.Config.Main).As(o.Config.Main),
		ErrSyncFailed,
	)
}

func (o Operation) commitChanges(message string) step {
	return func() error {
		o.Println("- Committing changes:", message)
		commit, err := o.Repository.CommitChanges(message)
		if err != nil {
			return errors.Wrap(err, ErrSyncFailed)
		}
		stats, err := commit.StatsContext(o.Context)
		if err == nil {
			o.Printf("-- Statistics:\n%s\n", stats)
		}
		return errors.Wrap(err, ErrSyncFailed)
	}
}
