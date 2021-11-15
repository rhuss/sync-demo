package config

// Project information.
type Project struct {
	Path       string
	ConfigPath string
}

// Projectlike is an interface for a project like object.
type Projectlike interface {
	GetPath() string
	GetConfigPath() string
}

func (p Project) GetPath() string {
	return p.Path
}

func (p Project) GetConfigPath() string {
	return p.ConfigPath
}
