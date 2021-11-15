package config

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func (c *Config) loadFromGit(project Projectlike) error {
	wr, ok := project.(withRepository)
	if !ok {
		return fmt.Errorf("%w: incompatible project", ErrConfigFileHaveInvalidFormat)
	}
	if c.Upstream == "" {
		c.Upstream = remoteURL(wr.Repository(), "upstream")
	}
	if c.Downstream == "" {
		c.Downstream = remoteURL(wr.Repository(), "downstream")
	}
	return nil
}

func remoteURL(repo *git.Repository, remoteName string) string {
	url := ""
	remote, err := repo.Remote(remoteName)
	if err == nil {
		url = remote.Config().URLs[0]
	}
	return url
}

type withRepository interface {
	Repository() *git.Repository
}
