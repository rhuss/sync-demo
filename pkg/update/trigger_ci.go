package update

import (
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"github.com/cardil/deviate/pkg/config/git"
	"github.com/cardil/deviate/pkg/errors"
)

func (o Operation) triggerCI() error {
	return triggerCI{o}.run()
}

func (o Operation) triggerCIMessage() string {
	return fmt.Sprintf(":robot: Triggering CI on branch `%s` after synching to `upstream/%s`",
		o.Config.Branches.ReleaseNext, o.Config.Branches.Main)
}

type triggerCI struct {
	Operation
}

func (c triggerCI) run() error {
	c.Println("Trigger CI")
	return runSteps([]step{
		c.checkout,
		c.addChange,
		c.commitChanges(c.triggerCIMessage()),
		c.pushBranch(c.Config.Branches.SynchCI),
	})
}

func (c triggerCI) checkout() error {
	remote := git.Remote{
		Name: "downstream",
		URL:  c.Config.Downstream,
	}
	err := c.Repository.Checkout(remote, c.Config.Branches.ReleaseNext).
		As(c.Config.Branches.SynchCI)
	return errors.Wrap(err, ErrUpdateFailed)
}

func (c triggerCI) addChange() error {
	filePath := path.Join(c.Project.Path, "ci")
	content := time.Now().Format(time.RFC3339)
	const fileReadableToOwnerPerm = 0o600
	err := ioutil.WriteFile(filePath, []byte(content), fileReadableToOwnerPerm)
	return errors.Wrap(err, ErrUpdateFailed)
}
