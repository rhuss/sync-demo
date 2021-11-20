package sync

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/cardil/deviate/pkg/errors"
)

func (o Operation) removeGithubWorkflows() error {
	o.Println("- Remove upstream Github workflows")
	workflows := path.Join(o.State.Project.Path, ".github", "workflows")

	dir, err := ioutil.ReadDir(workflows)
	if err != nil {
		return errors.Wrap(err, ErrSyncFailed)
	}
	for _, d := range dir {
		fp := path.Join(workflows, d.Name())
		if ok, _ := filepath.Match(o.GithubWorkflowsRemovalGlob, path.Base(fp)); ok {
			err = os.RemoveAll(fp)
			if err != nil {
				return errors.Wrap(err, ErrSyncFailed)
			}
		}
	}
	return nil
}
