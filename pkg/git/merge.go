package git

import (
	"fmt"

	"github.com/cardil/deviate/pkg/config/git"
	"github.com/cardil/deviate/pkg/errors"
	pkgfiles "github.com/cardil/deviate/pkg/files"
	gitv5 "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/magefile/mage/sh"
)

func (r Repository) Merge(remote *git.Remote, branch string) error {
	var (
		err    error
		before *plumbing.Reference
		after  *plumbing.Reference
	)
	if remote != nil {
		err = r.Fetch(*remote)
		if err != nil {
			return errors.Wrap(err, ErrRemoteOperationFailed)
		}
	}
	before, err = r.Head()
	if err != nil {
		return errors.Wrap(err, ErrLocalOperationFailed)
	}

	targetBranch := branch
	if remote != nil {
		targetBranch = fmt.Sprintf("%s/%s", remote.Name, branch)
	}
	// TODO: Consider rewriting this to Go native code.
	err = pkgfiles.WithinDirectory(r.Project.Path, func() error {
		return errors.Wrap(sh.RunV("git", "merge", "--commit",
			"--quiet", "--log", "-m", "Merge "+targetBranch, targetBranch),
			ErrRemoteOperationFailed)
	})
	if err != nil {
		_ = pkgfiles.WithinDirectory(r.Project.Path, func() error {
			return sh.RunV("git", "merge", "--abort")
		})
		return errors.Wrap(err, ErrRemoteOperationFailed)
	}
	after, err = r.Head()
	if err != nil {
		return errors.Wrap(err, ErrLocalOperationFailed)
	}

	if before.Hash() == after.Hash() {
		err = gitv5.NoErrAlreadyUpToDate
	}
	return err
}
