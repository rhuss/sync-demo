package git

import (
	"fmt"

	"github.com/cardil/deviate/pkg/config/git"
	"github.com/cardil/deviate/pkg/errors"
	gitv5 "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

func (r Repository) Push(remote git.Remote, branch string) error {
	repo := r.Repository
	specs := []config.RefSpec{
		refSpecForReferenceName(plumbing.NewBranchReferenceName(branch)),
	}
	auth, err := authentication(remote)
	if err != nil {
		return errors.Wrap(err, ErrLocalOperationFailed)
	}
	err = repo.PushContext(r.Context, &gitv5.PushOptions{
		RemoteName: remote.Name,
		RefSpecs:   specs,
		Auth:       auth,
		Force:      true,
	})
	return errors.Wrap(err, ErrRemoteOperationFailed)
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

func refSpecForReferenceName(name plumbing.ReferenceName) config.RefSpec {
	return config.RefSpec(
		fmt.Sprintf("%s:%s", name.String(), name.String()))
}
