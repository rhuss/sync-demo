package config

// New creates a new default configuration.
func New(project Projectlike) (Config, error) {
	c := newDefaults()
	err := c.load(project)
	if err != nil {
		return Config{}, err
	}
	err = c.validate()
	if err != nil {
		return Config{}, err
	}
	return c, nil
}

// newDefaults creates a new default configuration.
func newDefaults() Config {
	return Config{
		Branches: Branches{
			Main:            "main",
			ReleaseNext:     "release-next",
			ReleaseTemplate: "release-{{ .version.major }}-{{ .version.minor}}",
		},
	}
}
