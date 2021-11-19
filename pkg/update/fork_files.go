package update

import (
	"github.com/cardil/deviate/pkg/config/git"
	"github.com/cardil/deviate/pkg/errors"
)

func (o Operation) addForkFiles() error {
	return runSteps([]step{
		o.removeGithubWorkflows,
		func() error {
			o.Println("- Add fork's files")
			upstream := git.Remote{Name: "upstream", URL: o.Config.Upstream}
			err := o.Repository.Checkout(upstream, o.Config.Branches.Main).
				OntoWorkspace()
			return errors.Wrap(err, ErrUpdateFailed)
		},
		o.commitChanges(":open_file_folder: Update fork specific files"),
	})
}
