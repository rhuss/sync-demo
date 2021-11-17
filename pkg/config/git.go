package config

import (
	"github.com/cardil/deviate/pkg/config/git"
)

func (c *Config) loadFromGit(informer git.RemoteURLInformer) error {
	if c.Upstream == "" {
		c.Upstream = remoteURL(informer, "upstream")
	}
	if c.Downstream == "" {
		c.Downstream = remoteURL(informer, "downstream")
	}
	return nil
}

func remoteURL(informer git.RemoteURLInformer, remoteName string) string {
	url, _ := informer.Remote(remoteName)
	return url
}
