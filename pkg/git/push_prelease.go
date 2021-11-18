package git

import (
	"github.com/cardil/deviate/pkg/errors"
	"github.com/go-git/go-git/v5/plumbing"
)

func (r Repository) PushRelease(branch string) error {
	panic("implement me")
}

func (r Repository) DeleteBranch(branch string) error {
	err := r.Repository.DeleteBranch(branch)
	if err != nil {
		return errors.Wrap(err, ErrLocalOperationFailed)
	}
	ref := plumbing.NewBranchReferenceName(branch)
	err = r.Storer.RemoveReference(ref)
	return errors.Wrap(err, ErrLocalOperationFailed)
}
