package update

import (
	"github.com/cardil/deviate/pkg/state"
)

// Operation performs update - the upstream synchronization.
type Operation struct {
	state.State
}

func (o Operation) Run() error {
	steps := []step{
		o.mirrorBranches,
		o.resetReleaseNext,
		o.removeGithubWorkflows,
		o.addForkFiles,
		o.applyPatches,
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
