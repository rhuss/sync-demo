package git

import (
	"github.com/cardil/deviate/pkg/errors"
	gitv5 "github.com/go-git/go-git/v5"
)

func (r Repository) CommitChanges(message string) error {
	wt, err := r.Repository.Worktree()
	if err != nil {
		return errors.Wrap(err, ErrLocalOperationFailed)
	}
	err = wt.AddWithOptions(&gitv5.AddOptions{
		All:  true,
		Path: ".",
	})
	if err != nil {
		return errors.Wrap(err, ErrLocalOperationFailed)
	}
	_, err = wt.Commit(message, &gitv5.CommitOptions{
		All: true,
	})
	return errors.Wrap(err, ErrLocalOperationFailed)
}
