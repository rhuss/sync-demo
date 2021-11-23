package config

// Config for a deviate to operate.
type Config struct {
	Upstream                   string `yaml:"upstream" valid:"required"`
	Downstream                 string `yaml:"downstream" valid:"required"`
	DryRun                     bool   `yaml:"dryRun"`
	GithubWorkflowsRemovalGlob string `yaml:"githubWorkflowsRemovalGlob" valid:"required"`
	Branches                   `yaml:"branches"`
	Tags                       `yaml:"tags"`
}

// Tags holds configuration for tags.
type Tags struct {
	Synchronize bool   `yaml:"synchronize"`
	RefSpec     string `yaml:"refSpec" valid:"required"`
}

// Branches holds configuration for branches.
type Branches struct {
	Main             string `yaml:"main" valid:"required"`
	ReleaseNext      string `yaml:"releaseNext" valid:"required"`
	SynchCI          string `yaml:"synchCi" valid:"required"`
	ReleaseTemplates `yaml:"releaseTemplates"`
	Searches         `yaml:"searches"`
}

// ReleaseTemplates contains templates for release names.
type ReleaseTemplates struct {
	Upstream   string `yaml:"upstream" valid:"required"`
	Downstream string `yaml:"downstream" valid:"required"`
}

// Searches contains regular expressions used to search for branches.
type Searches struct {
	UpstreamReleases   string `yaml:"upstreamReleases" valid:"required"`
	DownstreamReleases string `yaml:"downstreamReleases" valid:"required"`
}
