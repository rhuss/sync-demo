package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/cardil/deviate/pkg/config/git"
	"gopkg.in/yaml.v3"
)

var (
	// ErrConfigFileCantBeRead when config file cannot be read.
	ErrConfigFileCantBeRead = errors.New("config file can't be read")
	// ErrConfigFileHaveInvalidFormat when config file has invalid format.
	ErrConfigFileHaveInvalidFormat = errors.New("config file have invalid format")
)

func (c *Config) load(project Project, informer git.RemoteURLInformer) error {
	bytes, err := os.ReadFile(project.ConfigPath)
	if err != nil {
		return fmt.Errorf("%s - %w: %v", project.ConfigPath,
			ErrConfigFileCantBeRead, err)
	}
	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		return fmt.Errorf("%s - %w: %v", project.ConfigPath,
			ErrConfigFileHaveInvalidFormat, err)
	}

	return c.loadFromGit(informer)
}
