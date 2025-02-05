package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/openshift-knative/deviate/pkg/errors"
)

func (c *Config) overrides() error {
	err := envconfig.Process("DEVIATE", c)
	return errors.Wrap(err, ErrConfigFileCantBeRead)
}
