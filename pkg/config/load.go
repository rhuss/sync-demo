package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	// ErrConfigFileCantBeRead when config file cannot be read.
	ErrConfigFileCantBeRead = errors.New("config file can't be read")
	// ErrConfigFileHaveInvalidFormat when config file has invalid format.
	ErrConfigFileHaveInvalidFormat = errors.New("config file have invalid format")
)

func (c *Config) load(project Projectlike) error {
	bytes, err := os.ReadFile(project.GetConfigPath())
	if err != nil {
		return fmt.Errorf("%s - %w: %v", project.GetConfigPath(),
			ErrConfigFileCantBeRead, err)
	}
	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		return fmt.Errorf("%s - %w: %v", project.GetConfigPath(),
			ErrConfigFileHaveInvalidFormat, err)
	}

	return c.loadFromGit(project)
}
