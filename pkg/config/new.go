package config

import "github.com/cardil/deviate/pkg/config/git"

// New creates a new default configuration.
func New(project Project, informer git.RemoteURLInformer) (Config, error) {
	c := newDefaults()
	err := c.load(project, informer)
	if err != nil {
		return Config{}, err
	}
	err = c.validate()
	if err != nil {
		return Config{}, err
	}
	return c, nil
}
