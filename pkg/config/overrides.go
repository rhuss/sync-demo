package config

import (
	"github.com/cardil/deviate/pkg/errors"
	"github.com/kelseyhightower/envconfig"
)

func (c *Config) overrides() error {
	err := envconfig.Process("DEVIATE", c)
	return errors.Wrap(err, ErrConfigFileCantBeRead)
}
