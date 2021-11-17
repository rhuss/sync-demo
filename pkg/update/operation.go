package update

import (
	"errors"

	"github.com/cardil/deviate/pkg/state"
)

// ErrUpdateFailed when the update failed.
var ErrUpdateFailed = errors.New("update failed")

// Operation performs update - the upstream synchronization.
type Operation struct {
	state.State
}

func (o Operation) Run() error {
	steps := []step{
		o.mirrorBranches,
		o.updateReleaseNext,
		o.triggerCI,
		o.createPR,
	}
	for _, s := range steps {
		if err := s(); err != nil {
			return err
		}
	}

	return nil
}

type step func() error
