package sync

import (
	"encoding/json"
	"fmt"

	"github.com/cardil/deviate/pkg/errors"
	"github.com/cardil/deviate/pkg/git"
	"github.com/cardil/deviate/pkg/github"
	"github.com/cardil/deviate/pkg/log/color"
)

func (o Operation) createPR() error {
	o.Println("Create a sync PR")
	pr := createPR{o}
	url, err := pr.active()
	if err != nil {
		if errors.Is(err, errPrNotFound) {
			return pr.open()
		}
		return err
	}

	o.Println("There is a PR open already at:", color.Yellow(*url))
	return nil
}

type createPR struct {
	Operation
}

var errPrNotFound = errors.New("PR not found")

func (c createPR) active() (*string, error) {
	repo, err := c.repository()
	if err != nil {
		return nil, errors.Wrap(err, ErrSyncFailed)
	}
	cl := github.NewClient(c.Project.Path,
		"pr", "list",
		"--repo", repo,
		"--state", "open",
		"--author", "@me",
		"--search", c.triggerCIMessage(),
		"--json", "url")
	cl.DisableColor = true
	cl.ProjectDir = c.Project.Path
	buff, err := cl.Execute(c.Context)
	if err != nil {
		return nil, errors.Wrap(err, ErrSyncFailed)
	}
	un := make([]map[string]interface{}, 0)
	err = json.Unmarshal(buff.Bytes(), &un)
	if err != nil {
		return nil, errors.Wrap(err, ErrSyncFailed)
	}

	if len(un) > 0 {
		u := fmt.Sprintf("%s", un[0]["url"])
		return &u, nil
	}
	return nil, errPrNotFound
}

func (c createPR) open() error {
	repo, err := c.repository()
	if err != nil {
		return errors.Wrap(err, ErrSyncFailed)
	}
	cl := github.NewClient(c.Project.Path,
		"pr", "create",
		"--repo", repo,
		"--body", fmt.Sprintf("This automated PR is to make sure the "+
			"forked project's `%s` branch (forked upstream's `%s` branch) passes"+
			" a CI.", c.Config.Branches.ReleaseNext, c.Config.Branches.Main),
		"--title", c.triggerCIMessage(),
		"--base", c.Config.Branches.ReleaseNext,
		"--head", c.Config.Branches.SynchCI)
	cl.ProjectDir = c.Project.Path
	buff, err := cl.Execute(c.Context)
	defer c.Println(buff.String())
	return errors.Wrap(err, ErrSyncFailed)
}

func (c createPR) repository() (string, error) {
	addr, err := git.ParseAddress(c.Config.Downstream)
	if err != nil {
		return "", errors.Wrap(err, ErrSyncFailed)
	}
	return addr.Path, nil
}
