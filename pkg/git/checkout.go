package git

import (
	"fmt"
	"io"

	"github.com/cardil/deviate/pkg/config/git"
	"github.com/cardil/deviate/pkg/errors"
	gitv5 "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

func (r Repository) Checkout(remote git.Remote, branch string) git.Checkout { //nolint:ireturn
	return &onGoingCheckout{
		remote: remote,
		branch: branch,
		repo:   r,
	}
}

type onGoingCheckout struct {
	remote git.Remote
	branch string
	repo   Repository
}

func (o onGoingCheckout) As(branch string) error {
	repo := o.repo.Repository
	hash, err := repo.ResolveRevision(o.revision())
	if err != nil {
		return errors.Wrap(err, ErrLocalOperationFailed)
	}
	wt, err := repo.Worktree()
	if err != nil {
		return errors.Wrap(err, ErrLocalOperationFailed)
	}
	var exist bool
	exist, err = o.branchExists(branch)
	if err != nil {
		return errors.Wrap(err, ErrLocalOperationFailed)
	}
	coOpts := &gitv5.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
	}
	if !exist {
		coOpts.Create = true
		coOpts.Hash = *hash

		err = repo.CreateBranch(&config.Branch{
			Name:   branch,
			Remote: o.remote.Name,
			Merge:  plumbing.NewBranchReferenceName(branch),
		})
		if err != nil {
			return errors.Wrap(err, ErrLocalOperationFailed)
		}
	}
	err = wt.Checkout(coOpts)
	if err != nil {
		return errors.Wrap(err, ErrLocalOperationFailed)
	}

	return errors.Wrap(wt.Reset(&gitv5.ResetOptions{
		Commit: *hash,
		Mode:   gitv5.HardReset,
	}), ErrLocalOperationFailed)
}

func (o onGoingCheckout) branchExists(branch string) (bool, error) {
	repo := o.repo.Repository
	iter, err := repo.Branches()
	if err != nil {
		return false, errors.Wrap(err, ErrLocalOperationFailed)
	}
	defer iter.Close()
	var ref *plumbing.Reference
	for ref, err = iter.Next(); !errors.Is(err, io.EOF); ref, err = iter.Next() {
		name := ref.Name()
		if name.IsBranch() && name.Short() == branch {
			return true, nil
		}
	}
	return false, nil
}

func (o onGoingCheckout) revision() plumbing.Revision {
	return plumbing.Revision(fmt.Sprintf("%s/%s", o.remote.Name, o.branch))
}
